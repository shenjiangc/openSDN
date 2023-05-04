package main

var (
	regions map[string]Region
	vpcs    map[string]Vpc
)

type Region struct {
	ID   string //`xml:"id" json:"id" description:"identifier of the region"`
	Name string //`xml:"name" json:"name" description:"name of the region" default:"beijing"`
}

type Vpc struct {
	ID   string //`xml:"id" json:"id" description:"identifier of the vpc"`
	Name string //`xml:"name" json:"name" description:"name of the vpc"`
}

func init() {
	regions = make(map[string]Region)
	vpcs = make(map[string]Vpc)
}
