//go:build linux
// +build linux

package core

/*
	#include <stdlib.h>
	#include <stdio.h>
	#include "healpix-utils.h"
	#include "healpix.h"
*/
import "C"

func (s *healpix) XyToOrder(xy int64) int64 {
	var hpixId int64
	if s.order == "ring" {
		hpixId = int64(C.healpixl_xy_to_ring(C.long(xy), C.int(s.nSide)))
	} else {
		hpixId = int64(C.healpixl_xy_to_nested(C.long(xy), C.int(s.nSide)))
	}
	return hpixId
}

func (s *healpix) RingToXy(ring_index int64) int64 {
	return int64(C.healpixl_ring_to_xy(C.long(ring_index), C.int(s.nSide)))
}

func (s *healpix) NestedToXY(nested_index int64) int64 {
	return int64(C.healpixl_nested_to_xy(C.long(nested_index), C.int(s.nSide)))
}

func (s *healpix) HealpixRangeConeSearch(ra float64, dec float64, radius float64, nSide int32, hpixIds []int64) []int64 {
	//var c_hpixIds = [256]C.longlong{} for darwin
	//var c_hpixIds = [256]C.long{} for linux
	// var c_hpixIds = [256]C.long{}
	var c_hpixIds = [256]C.long{}

	var iCount = C.healpix_rangesearch_radec_simple(C.double(ra), C.double(dec), C.double(radius), C.int(nSide), C.int(0), &c_hpixIds[0])
	var i int32

	hpixIds = make([]int64, iCount)
	for i = 0; i < int32(iCount); i++ {
		if s.order == "ring" {
			hpxid := int64(C.healpixl_xy_to_ring(c_hpixIds[i], C.int(nSide)))
			hpixIds[i] = hpxid
		} else {
			hpxid := int64(C.healpixl_xy_to_nested(c_hpixIds[i], C.int(nSide)))
			hpixIds[i] = hpxid
		}
	}
	return hpixIds
}

func (s *healpix) Neighbours(pixel_index int64) []int64 {
	var xy int64 = 0

	if s.order == "ring" {
		xy = s.RingToXy(pixel_index)
	} else {
		xy = s.NestedToXY(pixel_index)
	}

	var neighbours_c = [8]C.long{-1, -1, -1, -1, -1, -1, -1, -1}
	neighbours := make([]int64, 8)

	C.healpixl_get_neighbours(C.long(xy), &neighbours_c[0], C.int(s.nSide))
	for i := 0; i < 8; i++ {
		k := 4 - i
		if k < 0 {
			k += 8
		}
		xy_c := neighbours_c[k]
		var pixel int64 = -1
		if int64(xy) >= 0 {
			pixel = s.XyToOrder(int64(xy_c))
		}
		neighbours[k] = pixel
	}
	return neighbours
}
