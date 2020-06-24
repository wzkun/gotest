package main

import (
	"fmt"
)

type str1 struct {
	kind string
	tt   string
}
type str2 struct {
	*str1
	ne string
}

type str3 struct {
	*str1
	*str2
	nnn  string
	kind string
}

type int1 interface {
	ChangeKind(string)
	ChangeNe(string)
}

func (s *str1) ChangeKind(newstr string) {
	s.kind = newstr
}
func (s *str2) ChangeNe(newNe string) {
	s.ne = newNe
}
func (s *str3) ChangeKind(newstr string) {
	s.kind = newstr
}
func main() {
	var s1 str1 = str1{"kind", "tt"}
	var s2 str2 = str2{&s1, "ne"}
	var s3 str3 = str3{&s1, &s2, "nnn", "kind3"}

	var i1 int1 = &s3
	i1.ChangeKind("newkind")

	fmt.Println(s1.kind, s2.kind, s3.kind)

	
}
