package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	p2Extra()
	fmt.Println("time elapsed", time.Since(start).String())
	start = time.Now()
	p2()
	fmt.Println("time elapsed", time.Since(start).String())
}

type Coord struct {
	X, Y int
}

func (c Coord) Dist(other *Coord) int {
	dx := c.X - other.X
	if dx < 0 {
		dx *= -1
	}
	dy := c.Y - other.Y
	if dy < 0 {
		dy *= -1
	}
	return dx + dy
}

type Sensor struct {
	Location   *Coord
	Beacon     *Coord
	BeaconDist int
}

func lerp(a, b int, t float64) float64 {
	af := float64(a)
	bf := float64(b)
	return af + (bf-af)*t
}

func intersection(a, b, c, d *Coord) *Coord {
	dxcx := d.X - c.X
	aycy := a.Y - c.Y
	dycy := d.Y - c.Y
	axcx := a.X - c.X
	bxax := b.X - a.X
	byay := b.Y - a.Y

	// top := ((d.X - c.X) * (a.Y - c.Y)) - ((d.Y - c.Y) * (a.X - c.X))
	// bottom := ((b.X - a.X) * (d.Y - c.Y)) - ((b.Y - a.Y) * (d.X - c.X))
	top := (dxcx * aycy) - (dycy * axcx)
	bottom := (bxax * dycy) - (byay * dxcx)
	k := float64(top) / float64(bottom)
	if 0 > k || k > 1 || math.IsNaN(k) {
		return nil
	}
	x := lerp(a.X, b.X, k)
	y := lerp(a.Y, b.Y, k)
	o := &Coord{
		X: int(math.Ceil(x)),
		Y: int(math.Ceil(y)),
	}

	return o
}

type Data struct {
	Sensors   []*Sensor
	SensorMap map[int]map[int]struct{}
	BeaconMap map[int]map[int]struct{}
	MinX      int
	MaxX      int
	MinY      int
	MaxY      int
}

func (d *Data) addSensor(c *Coord) {
	if _, ok := d.SensorMap[c.X]; !ok {
		d.SensorMap[c.X] = make(map[int]struct{})
	}
	d.SensorMap[c.X][c.Y] = struct{}{}
}

func (d *Data) addBeacon(c *Coord) {
	if _, ok := d.BeaconMap[c.X]; !ok {
		d.BeaconMap[c.X] = make(map[int]struct{})
	}
	d.BeaconMap[c.X][c.Y] = struct{}{}
}

func (d *Data) IsBeacon(c *Coord) bool {
	if _, ok := d.BeaconMap[c.X]; !ok {
		return false
	}
	_, ok := d.BeaconMap[c.X][c.Y]
	return ok
}

func (d *Data) compareX(xs ...int) {
	for _, x := range xs {
		if x < d.MinX {
			d.MinX = x
		}
		if x > d.MaxX {
			d.MaxX = x
		}
	}
}

func (d *Data) compareY(ys ...int) {
	for _, y := range ys {
		if y < d.MinY {
			d.MinY = y
		}
		if y > d.MaxY {
			d.MaxY = y
		}
	}
}

func getData(test bool) (*Data, error) {
	filename := "input.txt"
	if test {
		filename = "test.txt"
	}
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	data := Data{
		MinX:      math.MaxInt,
		MaxX:      0,
		MinY:      math.MaxInt,
		MaxY:      0,
		SensorMap: make(map[int]map[int]struct{}),
		BeaconMap: make(map[int]map[int]struct{}),
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		sensorStr, beaconStr, _ := strings.Cut(scanner.Text(), ": ")
		sensorStr = strings.TrimPrefix(sensorStr, "Sensor at ")
		beaconStr = strings.TrimPrefix(beaconStr, "closest beacon is at ")
		sensor := getXY(sensorStr)
		beacon := getXY(beaconStr)
		data.compareX(sensor.X, beacon.X)
		data.compareY(sensor.Y, beacon.Y)
		data.Sensors = append(data.Sensors, &Sensor{Location: sensor, Beacon: beacon, BeaconDist: sensor.Dist(beacon)})
		data.addSensor(sensor)
		data.addBeacon(beacon)
	}

	return &data, nil
}

func getXY(s string) *Coord {
	xStr, yStr, _ := strings.Cut(s, ", ")
	return &Coord{X: intFromPair(xStr), Y: intFromPair(yStr)}
}

func intFromPair(s string) int {
	_, i, _ := strings.Cut(s, "=")
	v, _ := strconv.Atoi(i)
	return v
}

func p1() {
	data, err := getData(false)
	if err != nil {
		panic(err)
	}
	// y := 10
	y := 2_000_000
	var total int
	buffer := 25_000

	for x := data.MinX - buffer; x <= data.MaxX+buffer; x++ {
		coord := Coord{X: x, Y: y}
		if data.IsBeacon(&coord) {
			continue
		}
		for _, sensor := range data.Sensors {
			if sensor.Location.Dist(&coord) <= sensor.BeaconDist {
				total++
				break
			}
		}
	}
	fmt.Println(total)
}

type line struct {
	A *Coord
	B *Coord
}

func (l line) intersection(o *line) *Coord {
	return intersection(l.A, l.B, o.A, o.B)
}

func p2() {
	data, err := getData(false)
	if err != nil {
		panic(err)
	}
	var lines []*line
	for _, s := range data.Sensors {
		tp := &Coord{X: s.Location.X, Y: s.Location.Y - (s.BeaconDist + 1)}
		bp := &Coord{X: s.Location.X, Y: s.Location.Y + (s.BeaconDist + 1)}
		lp := &Coord{X: s.Location.X - (s.BeaconDist + 1), Y: s.Location.Y}
		rp := &Coord{X: s.Location.X + (s.BeaconDist + 1), Y: s.Location.Y}
		lines = append(lines, &line{A: tp, B: lp}, &line{A: lp, B: bp}, &line{A: bp, B: rp}, &line{A: rp, B: tp})
	}

	for i := 0; i < len(lines); i++ {
	jloop:
		for j := 0; j < len(lines); j++ {
			if i == j {
				// don't compare to itself
				continue
			}
			l1 := lines[i]
			l2 := lines[j]
			inter := l1.intersection(l2)
			if inter == nil {
				// lines don't intersect
				continue
			}
			if inter.X < 0 || inter.X > 4_000_000 {
				continue
			}
			if inter.Y < 0 || inter.Y > 4_000_000 {
				continue
			}
			for _, s := range data.Sensors {
				if s.Location.Dist(inter) <= s.BeaconDist {
					continue jloop
				}
			}
			fmt.Printf("%+v - %d\n", inter, (inter.X*4_000_000)+inter.Y)
			return
		}
	}
	fmt.Println("i did something wrong")
}

func p2Extra() {
	data, err := getData(false)
	if err != nil {
		panic(err)
	}
	for _, s := range data.Sensors {
		offset := s.BeaconDist + 1
		for j := 0; j <= offset; j++ {
			x1 := (s.Location.X + offset) - j
			x2 := (s.Location.X - offset) + j
			y1 := s.Location.Y - j
			y2 := s.Location.Y + j
			points := []*Coord{{X: x1, Y: y1}, {X: x1, Y: y2}, {X: x2, Y: y1}, {X: x2, Y: y2}}
			for _, p := range points {
				if p.X < 0 || p.X > 4_000_000 || p.Y < 0 || p.Y > 4_000_000 {
					continue
				}
				if check(p, s, data.Sensors) {
					fmt.Printf("%+v - %d\n", p, (p.X*4_000_000)+p.Y)
					return
				}
			}
		}
	}
}

func check(p *Coord, s *Sensor, sensors []*Sensor) bool {
	for _, sensor := range sensors {
		if sensor == s {
			continue
		}
		if sensor.Location.Dist(p) <= sensor.BeaconDist {
			return false
		}
	}
	return true
}
