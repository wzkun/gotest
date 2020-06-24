package main

import (
	"fmt"
)

type MyInt32 int32

func main() {
	var uid int32 = 12345
	var gid MyInt32 = (MyInt32)(uid)
	fmt.Printf("uid=%d, gid=%d\n", uid, gid)
}
