package main
import (
 	"fmt"
 	"github.com/wzkun/testmod"
 	testmodML "github.com/wzkun/testmod/v2"
)

func main() {
 	fmt.Println(testmod.Hi("Roberto"))
 	g, err := testmodML.Hi("Roberto", "pt")
 	if err != nil {
 		panic(err)
 	}
 	fmt.Println(g)
}