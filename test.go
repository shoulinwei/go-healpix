package main

import (
	"fmt"

	c "github.com/shoulinwei/go-healpix/core"
)

func main() {
	hp := c.NewHealpix(4, "ring")
	fmt.Printf("%d", hp.Npix())

	var hpxIds []int64

	hpxIds = hp.HealpixRangeConeSearch(192.0, 19, 5, 32, hpxIds)
	for _, v := range hpxIds {
		fmt.Printf("%d,", v)
	}

}
