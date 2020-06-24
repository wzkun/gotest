package main

import (
	"fmt"
)

type str2 struct {
	nnn  string
	kind string
	tt   string
}

type int1 interface {
	ChangeKind(string)
	ChangeNe(string)
}

func NewStr2(nnn, kind, tt string) *str2 {
	s := &str2{nnn, kind, tt}
	return s
}
func (s *str2) ChangeKind(newstr string) {
	s.kind = newstr
}
func (s *str2) ChangeNe(newstr string) {
	s.nnn = newstr
}

func main() {
	s := NewStr2("nnn", "kind", "tt")
	fmt.Println(s.kind, s.nnn, s.tt)
	s.ChangeKind("newKind")
	s.ChangeNe("newNNN")
	fmt.Println(s.kind, s.nnn, s.tt)
	var vi int1 = s
	vi.ChangeKind("kind222")
	vi.ChangeNe("nnn222")
	fmt.Println(s.kind, s.nnn, s.tt)

}
