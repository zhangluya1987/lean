package common

//var (
//	x1, y1 float64
//	x2, y2 float64
//	x3, y3 float64
//	x4, y4 float64
//)

func CalcSlope(x1, y1, x2, y2 float64) float64 {
	if x1 == x2 && y1 == y2 {
		return 0
	} else {
		lineSlope := (y2 - y1) / (x2 - x1)
		return float64(lineSlope)
	}

}
