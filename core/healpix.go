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
