package bmi

/*
1.BMI的计算结果 取其精度为2,后面小数点位数过多可能会引起计算结果异常。
2.BMI=体重(公斤)÷(身高×身高)(米)
*/

import (
	"strconv"
)

func GetBMI(Height, Weight float64) float64 {

	bmi := (Weight / (Height * Height))

	// 格式化浮点数并将浮点数转换为字符串,这里也可以采用fmt.Sprintf()
	bmiConvStr := strconv.FormatFloat(bmi, 'f', 2, 64)

	//ParseFloat 将字符串转换为浮点数
	bmiConvFloat, _ := strconv.ParseFloat(bmiConvStr, 64)
	return bmiConvFloat
}
