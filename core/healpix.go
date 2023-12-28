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

type healpix struct {
	nSide int32
	order string `default:"ring"` // ring or nested
}

func (s *healpix) Npix() int64 {
	return int64(12 * math.Pow(float64(s.nSide), 2))
}

func (s *healpix) PixelNsideValid(pixel int64) bool {
	return pixel >= 0 && pixel < s.Npix()
}

func (s *healpix) NestedToRing(nested_index int64) int64 {
	return int64(C.healpixl_nested_to_xy(C.long(nested_index), C.int(s.nSide)))
}

func (s *healpix) RingToNested(ring_index int64) int64 {
	return int64(C.healpixl_ring_to_xy(C.long(ring_index), C.int(s.nSide)))
}

func (s *healpix) RaDecDegToHealpix(ra float64, dec float64) int64 {
	hpixId := C.radec_to_healpixl(C.double(Deg2Rad(ra)), C.double(Deg2Rad(dec)), C.int(s.nSide))
	if s.order == "ring" {
		hpixId = C.healpixl_xy_to_ring(hpixId, C.int(s.nSide))
	} else {
		hpixId = C.healpixl_xy_to_nested(hpixId, C.int(s.nSide))
	}
	
	return int64(hpixId)
}

func NewHealpix(nSide int32, order string) *healpix {
	return &healpix{
		nSide: nSide,
		order: order,
	}
}
