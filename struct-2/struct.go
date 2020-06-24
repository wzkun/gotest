package main

import (
	"fmt"
)

type struct1 struct {
	kind string
	tt   string
}
type struct2 struct {
	name string
	id   string
	s1   struct1
}

func main() {
	var s struct2
	s.name = "name"
	s.id = "1233555"

	fmt.Println("s=", s)
	fmt.Println("kind=", s.s1.kind)
}
