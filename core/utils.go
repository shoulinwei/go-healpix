package core

const Pi = 3.1415926535897932384626
const RadPerDeg = 0.017453292519943295
const DEG_PER_RAD = 57.295779513082323

// pi / 180.
const RAD_PER_DEG = 0.017453292519943295

func Deg2Rad(deg float64) float64 {
	return deg * RAD_PER_DEG
}
