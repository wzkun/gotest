package main

import (
	"fmt"
)

// Region objects
type Region struct {
	OwerRegionType   string
	OwerRegionID     string
	OwerRegionCnName string
}

//Station objects
type Station struct {
	AlignmentID string  `json:"alignmentId,omitempty"`
	Number      float32 `json:"number"`
	Chain       string  `json:"chain,omitempty"`
}

// Bridge objects
type Bridge struct {
	ProjectID string
	ID        string

	Kind       string
	Type       string
	ModelType  string
	Region     Region   `json:"region,omitempty"`
	EndStation *Station `json:"endStation"`
}

func main() {
	cm := &Bridge{}
	cm.ProjectID = "dsfsfdf"
	cm.ID = "fsgergerdgedr"
	cm.Kind = "kind"
	cm.Type = "type"
	cm.ModelType = "modeltype"
	cm.Region.OwerRegionCnName = "OwerRegionCnName"
	cm.Region.OwerRegionID = "OwerRegionID"
	cm.Region.OwerRegionCnName = "OwerRegionCnName"

	cm.EndStation = &Station{}
	cm.EndStation.AlignmentID = "dsfsfsfs"

	fmt.Println("========cm======", cm)
	fmt.Println("========cm.Region======", cm.Region)
}
