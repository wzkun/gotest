package main

import (
	"fmt"
)

//Culvert struct
type Culvert struct {
	Name string
}

//UpdateName function
func (o *Culvert) UpdateName(name string) {
	o.Name = name
}

//CulvertPlan struct
type CulvertPlan struct {
	Name string
	Plan string
	Cul  Culvert
}

type CulvertPlanp struct {
	Name string
	Plan string
	Cul  *Culvert
}

func main() {
	culvert := Culvert{}
	culvert.Name = "dddd"
	culvertplan := CulvertPlan{}
	culvertplan.Cul = culvert

	culvertplan.Name = "abjg"
	fmt.Println("====culvertplan====", culvertplan)

	culvertplan.Cul.UpdateName("gggg")
	fmt.Println("====culvertplan====", culvertplan)

	culvertp := &Culvert{}
	culvertp.Name = "dddd"
	culvertplanp := &CulvertPlanp{}
	culvertplanp.Cul = culvertp

	culvertplanp.Name = "abjg"
	fmt.Println("====culvertplan====", culvertplanp.Cul)

	culvertplanp.Cul.UpdateName("gggg")
	fmt.Println("====culvertplan====", culvertplanp.Cul)

}
