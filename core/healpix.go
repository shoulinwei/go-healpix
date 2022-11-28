package core

/*
#cgo CFLAGS: -I../healpixlib
#cgo LDFLAGS: -L../healpixlib -lhpix
	#include <stdlib.h>
	#include <stdio.h>
	#include "healpix-utils.h"
	#include "healpix.h"
*/
import "C"
import (
	"math"
)

// service 服务
type healpix struct {
	nSide int32
	order string `default:"ring"`
}

func (s *healpix) Npix() int64 {
	return int64(12 * math.Pow(float64(s.nSide), 2))
}

func (s *healpix) HealpixRangeConeSearch(ra float64, dec float64, radius float64, nSide int32, hpixIds []int64) []int64 {
	//var c_hpixIds = [256]C.longlong{} for darwin
	//var c_hpixIds = [256]C.long{} for linux
	var c_hpixIds = [256]C.longlong{}
	var iCount = C.healpix_rangesearch_radec_simple(C.double(ra), C.double(dec), C.double(radius), C.int(nSide), C.int(0), &c_hpixIds[0])
	var i int32

	hpixIds = make([]int64, iCount)
	for i = 0; i < int32(iCount); i++ {
		hpxid := int64(C.healpixl_xy_to_nested(c_hpixIds[i], C.int(nSide)))
		hpixIds[i] = hpxid
	}
	return hpixIds
}
func (s *healpix) RaDecDegToHealpix(ra float64, dec float64, nSide int32) int32 {
	var hpixId = C.radec_to_healpixl(C.double(ra), C.double(dec), C.int(nSide))
	return int32(hpixId)
}

func NewHealpix(nSide int32, order string) *healpix {
	return &healpix{
		nSide: nSide,
		order: order,
	}
}
