package main

import (
	"fmt"
	"time"
)

// TimeFromMillis function
func TimeFromMillis(millis int64) time.Time {
	return time.Unix(0, millis*int64(time.Millisecond))
}
func main() {
	tim := int64(1542333051747)
	times := time.Unix(tim, 0)
	fmt.Println("==========times=====", times)

	times2 := TimeFromMillis(1542172145405)
	fmt.Println("==========times2=====", times2)
	times2 = TimeFromMillis(1542333051747)
	fmt.Println("==========times2=====", times2)
	times3 := time.Now().String()
	fmt.Println("=====times3====", times3)
}
