package main

import (
	"fmt"
	"os"
	"strings"
)

var fname = "assets/input.txt"

func main() {
	P2()
}

func P2() {
	s := parse(fname)

	svr := s.GetOrCreateDevice("svr")
	out := s.GetOrCreateDevice("out")
	dac := s.GetOrCreateDevice("dac")
	fft := s.GetOrCreateDevice("fft")

	dac_fft := fromTo(dac, fft)
	fmt.Printf("dac_fft: %v\n", dac_fft)

	fft_dac := fromTo(fft, dac)
	fmt.Printf("fft_dac: %v\n", fft_dac)

	svr_dac := fromTo(svr, dac)
	fmt.Printf("svr_dac: %v\n", svr_dac)

	svr_fft := fromTo(svr, fft)
	fmt.Printf("svr_fft: %v\n", svr_fft)

	fft_out := fromTo(fft, out)
	fmt.Printf("fft_out: %v\n", fft_out)

	dac_out := fromTo(dac, out)
	fmt.Printf("dac_out: %v\n", dac_out)

	total := (svr_dac * dac_fft * fft_out) +
		(svr_fft * fft_dac * dac_out)
	fmt.Printf("total: %v\n", total)
}

func P1() {
	s := parse(fname)
	you := s.GetOrCreateDevice("you")
	out := s.GetOrCreateDevice("out")
	fmt.Printf("fromTo(you, out): %v\n", fromTo(you, out))
}

func fromTo(from, to *Device) int {
	cache := make(map[string]int)
	return fromToMemo(from, to, cache)
}

func fromToMemo(from, to *Device, cache map[string]int) int {
	// Create cache key from device identifiers
	cacheKey := from.Ident + "->" + to.Ident

	// Check cache first
	if result, exists := cache[cacheKey]; exists {
		return result
	}

	if from == to {
		return 1
	}

	var count int
	for _, output := range from.Outputs {
		count += fromToMemo(output, to, cache)
	}
	// Cache the result before returning
	cache[cacheKey] = count
	return count
}

type Device struct {
	Ident   string
	Inputs  []*Device
	Outputs []*Device
}

func (d *Device) AddOutput(other *Device) {
	d.Outputs = append(d.Outputs, other)
	other.Inputs = append(other.Inputs, d)
}

type Servers struct {
	All map[string]*Device
}

func (s *Servers) GetOrCreateDevice(ident string) *Device {
	device := s.All[ident]
	if device == nil {
		device = &Device{Ident: ident}
		s.All[ident] = device
	}
	return device
}

func parse(fname string) *Servers {
	data, err := os.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	out := Servers{All: make(map[string]*Device)}
	for _, line := range lines {
		ident, parts, _ := strings.Cut(line, ": ")
		device := out.GetOrCreateDevice(ident)
		for _, p := range strings.Split(strings.TrimSpace(parts), " ") {
			downstream := out.GetOrCreateDevice(p)
			device.AddOutput(downstream)
		}
	}

	return &out
}
