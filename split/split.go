package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "371f6563-556c-4c36-b175-448941e5d6da,371f6563-556c-4c36-b175-448941e5d6db"
	a := strings.Split(str, ",")

	fmt.Println(len(str))

	fmt.Println(len(a))
	fmt.Println(a)

	str = ""
	if str == "" {
		fmt.Println(len(str))
	} else {
		a = strings.Split(str, ",")

		fmt.Println(len(a))
		fmt.Println(a)
	}

}
