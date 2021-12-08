package main

import (
	"fmt"
	"lesson1/homework/whetherParallel/check"
	"lesson1/homework/whetherParallel/common"
	"lesson1/homework/whetherParallel/input"
)



//注意 数据类型如果定义为int类型 四舍五入的计算方式 会影响结果的准确性。
func main(){
	input.Input()
	line1 := common.CalcSlope(input.X1,input.Y1,input.X2,input.Y2)
	line2 := common.CalcSlope(input.X3,input.Y3,input.X4,input.Y4)

	calcResult := check.CheckParallel(line1,line2)

	if calcResult == true {
		fmt.Printf("line = %v,line2=%v,两条直线斜率相等，平行\n",line1,line2)
	}else{
		fmt.Printf("line = %v,line2=%v,两条直线斜率不相等，不平行\n",line1,line2)
	}
}