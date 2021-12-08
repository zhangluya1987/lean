package common

import "strconv"

func StrToFloat64(str string) float64 {
	float64Result,_ := strconv.ParseFloat(str,64)
	return float64Result
}

func Float64ToStr(x float64) string {
	strResult := strconv.FormatFloat(x,'f',2,64)
	return strResult
}
func AvgBFR(x []float64) float64 {
	var total float64
	for i:=0;i<len(x);i++{
		total+=x[i]
	}
	return total/float64(len(x))
}
