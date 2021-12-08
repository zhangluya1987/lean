package sex

import (
	"fmt"
)

func GetSexWeight(Sex string) int {
	//(男为1，女为0)
	var sexWeight int
	if Sex == "男" {
		sexWeight = 1
	} else if Sex == "女"{
		sexWeight = 0
	}else{
		fmt.Println("输入错误,请重新输入")
	}
	return sexWeight
}