package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	a := Almanac{}
	if err := a.Load("assets/input.txt"); err != nil {
		panic(err)
	}
	fmt.Println(p1(&a))
	fmt.Println(p2(&a))
}

func p1(a *Almanac) int {
	lowest := math.MaxInt
	for _, seed := range a.Seeds {
		soil := a.SoilMap.Convert(seed)
		fert := a.FertilizerMap.Convert(soil)
		water := a.WaterMap.Convert(fert)
		light := a.LightMap.Convert(water)
		temp := a.TempMap.Convert(light)
		hum := a.HumidityMap.Convert(temp)
		location := a.LocationMap.Convert(hum)
		if location < lowest {
			lowest = location
		}
	}
	return lowest
}

func p2(a *Almanac) int {
	for _, rng := range a.LocationMap.Ranges {
		for i := 0; i < rng.rng; i++ {
			location := rng.destStart + i
			hum := a.LocationMap.Invert(location)
			temp := a.HumidityMap.Invert(hum)
			light := a.TempMap.Invert(temp)
			water := a.LightMap.Invert(light)
			fert := a.WaterMap.Invert(water)
			soil := a.FertilizerMap.Invert(fert)
			seed := a.SoilMap.Invert(soil)
			for _, seedrng := range a.SeedRanges {
				if seedrng.Contains(seed) {
					return location
				}
			}
		}
	}
	return -1
}

type Range struct {
	srcStart  int
	destStart int
	rng       int
}

func (r Range) Convert(src int) int {
	if offset := src - r.srcStart; offset >= 0 && offset < r.rng {
		return r.destStart + offset
	}
	return -1
}

func (r Range) Invert(dest int) int {
	if offset := dest - r.destStart; offset >= 0 && offset < r.rng {
		return r.srcStart + offset
	}
	return -1
}

type Mapping struct {
	Ranges []*Range
}

func (m *Mapping) Load(data string) {
	parts := strings.Split(data, "\n")
	for _, part := range parts[1:] {
		if part == "" {
			continue
		}
		line := AtoIs(part)
		m.Ranges = append(m.Ranges, &Range{
			destStart: line[0],
			srcStart:  line[1],
			rng:       line[2],
		})
	}
	sort.Slice(m.Ranges, func(i, j int) bool {
		return m.Ranges[i].destStart < m.Ranges[j].destStart
	})
}

func (m Mapping) Convert(src int) int {
	for _, r := range m.Ranges {
		if dest := r.Convert(src); dest != -1 {
			return dest
		}
	}
	return src
}

func (m Mapping) Invert(dest int) int {
	for _, r := range m.Ranges {
		if src := r.Invert(dest); src != -1 {
			return src
		}
	}
	return dest
}

type SeedRange struct {
	start int
	rng   int
}

func (s *SeedRange) Contains(seed int) bool {
	offset := seed - s.start
	return offset >= 0 && offset < s.rng
}

type Almanac struct {
	Seeds         []int
	SeedRanges    []*SeedRange
	SoilMap       Mapping
	FertilizerMap Mapping
	WaterMap      Mapping
	LightMap      Mapping
	TempMap       Mapping
	HumidityMap   Mapping
	LocationMap   Mapping
}

func (a *Almanac) Load(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	data, err := ReadAll(f)
	if err != nil {
		return err
	}

	parts := strings.Split(data, "\n\n")

	a.Seeds = AtoIs(strings.TrimPrefix(parts[0], "seeds: "))
	for i := 1; i < len(a.Seeds); i += 2 {
		start := a.Seeds[i-1]
		rng := a.Seeds[i]
		a.SeedRanges = append(a.SeedRanges, &SeedRange{start: start, rng: rng})
	}

	a.SoilMap.Load(parts[1])
	a.FertilizerMap.Load(parts[2])
	a.WaterMap.Load(parts[3])
	a.LightMap.Load(parts[4])
	a.TempMap.Load(parts[5])
	a.HumidityMap.Load(parts[6])
	a.LocationMap.Load(parts[7])

	return nil
}

func ReadAll(r io.Reader) (string, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func AtoIs(s string) []int {
	bNums := strings.Split(s, " ")
	out := make([]int, len(bNums))
	for i, bNum := range bNums {
		out[i], _ = strconv.Atoi(string(bNum))
	}
	return out
}
