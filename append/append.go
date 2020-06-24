package main

import (
	"fmt"
	"time"
)

type ssmm struct {
	jk int64
}

func main() {
	ss := make([]int, 10)

	ss = append(ss, 1)
	fmt.Println("===========ss=======", ss)

	mm := make([]*ssmm, 0, 1)

	ms := &ssmm{
		jk: 123,
	}

	mm = append(mm, ms)

	fmt.Println("===============mm==========", mm)

	tims := time.Now().Unix()
	fmt.Println("===========tims========", tims)
	tm := time.Unix(tims, 0)
	fmt.Println("===========tm========", tm)
}
