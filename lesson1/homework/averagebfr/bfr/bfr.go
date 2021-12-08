package bfr

import (
	"strconv"
)

func GetBFR(BMI float64, Age float64, SexWeight int) float64 {

	//体脂率:1.2×BMI+0.23×年龄-5.4-10.8×性别(男为1，女为0) 这里换算成百分比进行计算
	bfr := ((1.2 * BMI) + (0.23 * Age) - 5.4 - (10.8 * float64(SexWeight))) / 100

	bfrConvStr := strconv.FormatFloat(bfr, 'f', 2, 64)

	bfrConvFloat, _ := strconv.ParseFloat(bfrConvStr, 64)
	return bfrConvFloat
}

func GetBFRAverage(x []float64) float64 {
	var Sum float64
	for index,_ := range x{
		Sum +=x[index]
	}
	return Sum/float64(len(x))
}
