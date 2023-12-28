package main

import (
	"fmt"

	c "github.com/shoulinwei/go-healpix/core"
)

func main() {
	//---------------init-------------------
	hp := c.NewHealpix(256, "nested")

	//---------------number of pixels-------------------
	fmt.Printf("total pixel:%d\n", hp.Npix())

	//---------------get Healpix Id via RA,DEC in Degree-------------------
	hpxId := hp.RaDecDegToHealpix(246.37314410833565,39.72582110537353) 
	fmt.Printf("%v\n", hpxId) //151479
	hpxId = hp.RaDecDegToHealpix(246.43614768338378,39.87359713842207)
	fmt.Printf("%v\n", hpxId) //151528

	//---------------get Healpix Ids via Cone Search-------------------
	var hpxIds []int64
	hpxIds = hp.HealpixRangeConeSearch(192.0, 19, 5, 32, hpxIds)
	for _, v := range hpxIds {
		fmt.Printf("%d,", v)
	}
	fmt.Printf("\n")

	//---------------get 8 neighbours-------------------
	neighbours := hp.Neighbours(151479)
	for _, v := range neighbours {
		fmt.Printf("%d,", v)
	}
	fmt.Printf("\n")

}
