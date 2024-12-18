package main

func main() {
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
