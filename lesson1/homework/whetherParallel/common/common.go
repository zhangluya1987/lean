package common

var(
	x1 int
	y1 int
	x2 int
	y2 int
	x3 int
	y3 int
	x4 int
	y4 int
)

func CalcSlope(x1,y1,x2,y2 int) float64 {
	if x1 == x2 && y1 == y2 {
		return 0
	} else {
		lineSlope := (y2 - y1) / (x2 - x1)
		return float64(lineSlope)
	}

}

