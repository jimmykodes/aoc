package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"slices"
)

func main() {
	data, err := os.ReadFile("assets/input.txt")
	if err != nil {
		panic(err)
	}
	var m Machine
	if err := json.Unmarshal(data, &m); err != nil {
		panic(err)
	}

	m.A = 0o4354337671237000
	incr := 0o1
	// for j := 0; j < 10; j++ {
	for {
		// i++
		out := m.Run()
		// if len(out) > len(m.Instructions) {
		// 	fmt.Println("A too high, instructions too long")
		// 	return
		// }
		//
		fmt.Printf("%o, %v, %d, %s\n", m.A, out, len(out), "2414754114550330")
		if slices.Equal(out, m.Instructions) {
			fmt.Println(m.A)
			return
		}
		m.A += incr
	}
}

type Machine struct {
	ip           int
	A, B, C      int
	Instructions []int
	Buffer       []int
}

func (m Machine) Run() []int {
	for {
		if m.ip >= len(m.Instructions) {
			return m.Buffer
		}
		op := OpCode(m.Instructions[m.ip])
		if op.Run(m.Instructions[m.ip+1], &m) {
			m.ip += 2
		}
	}
}

type OpCode int

const (
	adv OpCode = iota
	bxl
	bst
	jnz
	bxc
	out
	bdv
	cdv
)

func (o OpCode) String() string {
	return [...]string{
		adv: "adv",
		bxl: "bxl",
		bst: "bst",
		jnz: "jnz",
		bxc: "bxc",
		out: "out",
		bdv: "bdv",
		cdv: "cdv",
	}[o]
}

func (o OpCode) Run(operand int, machine *Machine) bool {
	return operations[o](operand, machine)
}

func combo(operand int, machine *Machine) int {
	switch operand {
	case 4:
		return machine.A
	case 5:
		return machine.B
	case 6:
		return machine.C
	case 7:
		panic("invalid operand")
	default:
		return operand
	}
}

var operations = [...]func(operand int, machine *Machine) bool{
	adv: func(operand int, machine *Machine) bool {
		machine.A /= int(math.Pow(2, float64(combo(operand, machine))))
		return true
	},
	bxl: func(operand int, machine *Machine) bool {
		machine.B ^= operand
		return true
	},
	bst: func(operand int, machine *Machine) bool {
		machine.B = combo(operand, machine) % 8
		return true
	},
	jnz: func(operand int, machine *Machine) bool {
		if machine.A == 0 {
			return true
		}
		machine.ip = operand
		return false
	},
	bxc: func(operand int, machine *Machine) bool {
		machine.B ^= machine.C
		return true
	},
	out: func(operand int, machine *Machine) bool {
		v := combo(operand, machine) % 8
		machine.Buffer = append(machine.Buffer, v)
		return true
	},
	bdv: func(operand int, machine *Machine) bool {
		machine.B = machine.A / int(math.Pow(2, float64(combo(operand, machine))))
		return true
	},
	cdv: func(operand int, machine *Machine) bool {
		machine.C = machine.A / int(math.Pow(2, float64(combo(operand, machine))))
		return true
	},
}
