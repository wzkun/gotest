package main

import (
	"encoding/json"
	"fmt"
	"sort"

	"zinto.db/golang/decode"
)

// TunnelSum objects
type TunnelSum struct {
	ProjectID string
	ID        string

	SectionID string
	Name      string
	EnName    string

	TotalGroups     TunnelGroup
	CompletedGroups TunnelGroup

	TotalPartCount          map[string]float32
	TotalCompletedPartCount map[string]float32
}

// MinUint64 function
func MinUint64(a uint64, b uint64) uint64 {
	if a > b {
		return b
	}
	return a
}

// MaxUint64 function
func MaxUint64(a uint64, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}

// Interval objects
type Interval struct {
	Start uint64
	End   uint64
}

// NewInterval function
func NewInterval(start, end uint64) *Interval {
	iv := &Interval{
		Start: start,
		End:   end,
	}
	return iv
}

// BeforeInterval function
func (o *Interval) BeforeInterval(iv *Interval) bool {
	return o.End < iv.Start
}

// Contain function
func (o *Interval) Contain(number uint64) bool {
	return (o.Start <= number) && (number <= o.End)
}

// AfterInterval function
func (o *Interval) AfterInterval(iv *Interval) bool {
	return o.Start > iv.End
}

//IntervalList  type
type IntervalList []*Interval

// Sort function
func (o *IntervalList) Sort() {
	ivs := *o
	sort.Slice(ivs, func(i, j int) bool {
		return ivs[i].Start < ivs[j].End
	})
}

// SortedIntervalSet objects
type SortedIntervalSet struct {
	ivs IntervalList
}

// AddInterval function
func (o *SortedIntervalSet) AddInterval(start, end uint64) {
	civ := NewInterval(start, end)
	before := make(IntervalList, 0, len(o.ivs))
	after := make(IntervalList, 0, len(o.ivs))
	res := make(IntervalList, 0, len(o.ivs))

	for _, iv := range o.ivs {
		if iv.BeforeInterval(civ) {
			before = append(before, iv)
		}
		if iv.Contain(civ.Start) || iv.Contain(civ.End) {
			civ.Start = MinUint64(civ.Start, iv.Start)
			civ.End = MaxUint64(civ.End, iv.End)
		}
		if iv.AfterInterval(civ) {
			after = append(after, iv)
		}
	}

	res = append(res, before...)
	res = append(res, civ)
	res = append(res, after...)

	res.Sort()
	o.ivs = res
}

// TunnelUint objects
type TunnelUint struct {
	*SortedIntervalSet
	inited bool

	GroupType   string
	GroupNumber string
	ModelType   string
	Direction   string
	Intervals   string
}

// NewSortedIntervalSet function
func NewSortedIntervalSet() *SortedIntervalSet {
	o := &SortedIntervalSet{}
	o.ivs = make(IntervalList, 0, 100)
	return o
}

// NewSortedIntervalSetFromJSONString function
func NewSortedIntervalSetFromJSONString(data string) (*SortedIntervalSet, error) {
	o := NewSortedIntervalSet()
	err := decode.JSON.Unmarshal(json.RawMessage(data), &o.ivs)
	return o, err
}

// ToJSONString function
func (o *SortedIntervalSet) ToJSONString() string {
	js, _ := decode.JSON.Marshal(o.ivs)
	return string(js)
}

// NewTunnelUnit function
func NewTunnelUnit() *TunnelUint {
	o := &TunnelUint{}
	return o
}
func (o *TunnelUint) checkInit() {
	if o.inited == false {
		o.SortedIntervalSet, _ = NewSortedIntervalSetFromJSONString(o.Intervals)
		o.inited = true
	}
}

// AddInterval function
func (o *TunnelUint) AddInterval(start, end float32) {
	o.checkInit()

	st := uint64(start)
	ed := uint64(end)
	o.SortedIntervalSet.AddInterval(st, ed)
	o.Intervals = o.SortedIntervalSet.ToJSONString()
}

// TunnelGroup objects
type TunnelGroup struct {
	inited bool
	proxy  map[string]*TunnelUint

	Values string
}

// getTunnelUnit function
func (o *TunnelGroup) getTunnelUnit(groupType, groupNumber, modelType, direction string) *TunnelUint {
	o.checkInit()

	key := modelType + groupType + groupNumber + direction
	if u, hit := o.proxy[key]; hit {
		return u
	}

	u := NewTunnelUnit()
	u.ModelType = modelType
	u.GroupType = groupType
	u.GroupNumber = groupNumber
	u.Direction = direction
	o.proxy[key] = u
	return u
}

func (o *TunnelGroup) checkInit() {
	if o.inited == false {
		o.proxy = map[string]*TunnelUint{}
		decode.JSON.UnmarshalFromString(o.Values, &o.proxy)
		o.inited = true
	}
}

// AddInterval function
func (o *TunnelGroup) AddInterval(groupType, groupNumber, modelType, direction string, start, end float32) {
	unit := o.getTunnelUnit(groupType, groupNumber, modelType, direction)
	unit.AddInterval(start, end)
}

// AddPart function
func (o *TunnelSum) AddPart(groupType, groupNumber, direction, modelType string, start, end float32) {
	o.TotalPartCount[modelType]++
	o.TotalGroups.AddInterval(groupType, groupNumber, modelType, direction, start, end)
}

// TunnelPlanSum objects
type TunnelPlanSum struct {
	*TunnelSum
	PlanName string `json:"planName,omitempty"`
}

func main() {
	tunnelsum := &TunnelSum{}
	tunnelsum.TotalPartCount = map[string]float32{}
	tunnelsum.TotalCompletedPartCount = map[string]float32{}
	tunnelplansum := &TunnelPlanSum{}
	tunnelplansum.TunnelSum = tunnelsum
	tunnelplansum.AddPart("ZhuDong", "0", "Left", "GongQiangChuPen", float32(31834), float32(31835))

	fmt.Println("tunnelsum=", tunnelsum)
	fmt.Println("tunnelplansum.TotalGroups=", tunnelplansum.TotalGroups)
}
