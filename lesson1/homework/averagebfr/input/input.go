package input

import (
	"fmt"
	"strconv"
	"strings"
)

//func InputData(i){
//
//	//println("会主动换行"),print(不会主动换行)
//	fmt.Println("格式: 姓名:Jesse 性别(男/女) 身高(m) 体重(kg) 年龄")
//	fmt.Scanln("&Name, &Sex, &Height, &Weight, &Age)
//}

var (
	Name       string
	Sex        string
	Height     float64
	Weight     float64
	Age        float64
	PeopleInfo [][]string = make([][]string, 3)
)

func Input(){
	fmt.Printf("输入格式(中间空格隔开): 姓名 性别(男/女) 身高(m) 体重(kg) 年龄\n")
	for i := range PeopleInfo {
		fmt.Scanln(&Name, &Sex, &Height, &Weight, &Age)

		strHeight := strconv.FormatFloat(Height, 'f', 2, 64)
		strWeight := strconv.FormatFloat(Weight, 'f', 2, 64)
		strAge := strconv.FormatFloat(Age, 'f', 2, 64)
		PeopleInfo[i] = []string{Name, Sex, strHeight, strWeight, strAge}
	}
	fmt.Println(strings.Repeat("-",80))
}