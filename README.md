

go-healpix
=========
go-healpix is a Go wrapper around HEALPix C library

# Install
    go get github.com/shoulinwei/go-healpix
	git clone https://github.com/shoulinwei/go-healpix.git
	cd go-healpix/healpixlib

	//at MAC
	make dylib
	//at linux
	make so

	copy libhpix.dylib or libhpix.so to your project root

# Usage
```go
    import (
        "fmt"
        c "github.com/shoulinwei/go-healpix/core"
    )
    //---------------init-------------------
	hp := c.NewHealpix(256, "nested")

	//---------------number of pixels-------------------
	fmt.Printf("total pixel:%d\n", hp.Npix())

	//---------------get Healpix Id via RA,DEC in Degree-------------------
	hpxId := hp.RaDecDegToHealpix(246.37314410833565,39.72582110537353) 
	fmt.Printf("%v\n", hpxId) //151479
	hpxId = hp.RaDecDegToHealpix(246.43614768338378,39.87359713842207)
	fmt.Printf("%v\n", hpxId) //151528

	//---------------get Healpix Ids covered via center's RA, DEC and radius------
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
```    