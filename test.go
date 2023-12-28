package main

import (
	"fmt"

	c "github.com/shoulinwei/go-healpix/core"
)

func main() {
	hp := c.NewHealpix(256, "nested")
	fmt.Printf("total pixel:%d\n", hp.Npix())

	hpxId := hp.RaDecDegToHealpix(246.37314410833565,39.72582110537353) 
	fmt.Printf("%v\n", hpxId) //151479
	hpxId = hp.RaDecDegToHealpix(246.43614768338378,39.87359713842207)
	fmt.Printf("%v\n", hpxId) //151528

	var hpxIds []int64
	hpxIds = hp.HealpixRangeConeSearch(192.0, 19, 5, 32, hpxIds)
	for _, v := range hpxIds {
		fmt.Printf("%d,", v)
	}
	fmt.Printf("\n")
}
