package main

import (
	"fmt"
)

type Human struct {
	name   string
	height string
	weight string
}

type Action struct {
	Human
}

type HumanAction interface {
	beHappy()
}

func (h *Human) beHappy() {
	fmt.Printf("%s is so happy", h.name)
}

func main() {
	person := &Action{
		Human{name: "Vova",
			height: "185",
			weight: "80",
		},
	}

	person.beHappy()
}
