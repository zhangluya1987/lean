package main

import (
	"fmt"
	"lesson1/homework/averagebfr/bfr"
	"lesson1/homework/averagebfr/bmi"
	"lesson1/homework/averagebfr/common"
	"lesson1/homework/averagebfr/input"
	"lesson1/homework/averagebfr/sex"
	"lesson1/homework/averagebfr/suggest"
)


func main() {
	var peopleSlice []string
	var bfrSlice []float64

	//执行数据输入操作,会将结果直接传入一个二维slice
	input.Input()

	//计算每个人的bmi 这里从input.Input()里面提取输入信息
	for i := range input.PeopleInfo{
		getName := input.PeopleInfo[i][0]
		getSex := input.PeopleInfo[i][1]
		getAge := common.StrToFloat64(input.PeopleInfo[i][4])

		newHeight := common.StrToFloat64(input.PeopleInfo[i][2])
		newWeight := common.StrToFloat64(input.PeopleInfo[i][3])
		calc_sex_weight := sex.GetSexWeight(input.Sex)
		calc_bmi := bmi.GetBMI(newHeight,newWeight)
		calc_bfr := bfr.GetBFR(calc_bmi,getAge,calc_sex_weight)
		getSuggest := suggest.Suggest(getSex,getAge,calc_bfr)

		//将获取到的人员姓名加入slice
		peopleSlice = append(peopleSlice,getName)
		bfrSlice = append(bfrSlice,calc_bfr)
		fmt.Printf("姓名:%v BMI:%v 体脂率:%v,建议:%v\n",getName,calc_bmi,calc_bfr,getSuggest)
	}

	avgBFR := common.AvgBFR(bfrSlice)
	totalPeople := len(peopleSlice)
	fmt.Printf("总人数:%v,平均体脂率:%v\n",totalPeople,avgBFR)
}