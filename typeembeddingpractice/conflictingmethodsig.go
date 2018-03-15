package main

import (
	"fmt"
)

type Ball struct {
	Radius   int
	Material string
}

type Football struct {
	Ball
	Bounce int
}

func (b Ball) Bounce() {
	fmt.Printf("Radius = %d\n", b.Radius)
}

func main() {
	fb := Football{Ball{Radius: 5, Material: "leather"}, 5}
	fb.Bounce()
	fmt.Printf("b = %+v\n", fb)
}
