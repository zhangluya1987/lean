package bfr

import (
	"fmt"
	"lesson3/homework/02.fatRate_UnitTesting_v1.0/bmi"
	"lesson3/homework/02.fatRate_UnitTesting_v1.0/suggest"
	"testing"
)

// 1.录入正常 BMI、年龄、性别，确保计算结果符合预期；
func TestCalcFatRateCase1(t *testing.T) {
	weightKG,heightM := 70.0,1.75
	age,sex := 23,"男"
	bmiResult,err:= bmi.CalcBMI(weightKG,heightM)

	if err != nil {
		t.Errorf("%v",err)
	}

	bfrResult,err:= CalcFatRate(bmiResult,age,sex)
	if err != nil{
		fmt.Errorf("%v",err)
	}

	if bfrResult != 0.1651857142857143 {
		t.Errorf("预期计算结果是0.1651857142857143,实际计算结果是%v\n",bfrResult)
	}
}

// 2.录入非法 BMI、年龄、性别（0、负数、超过 150 的年龄、非男女的性别输入），返回错误；
func TestCalcFatRateCase2(t *testing.T) {

	// 录入非法bmi=0进行测试
	{
		age, sex := 23, "男"
		bmiResult := 0
		_,err := CalcFatRate(float64(bmiResult), age, sex)
		if err != nil {
			t.Errorf("%v\n", err)
		}
	}

	// 录入非法bmi=-1进行测试
	{
		age, sex := 23, "男"
		bmiResult := -1
		_,err := CalcFatRate(float64(bmiResult), age, sex)
		if err != nil {
			t.Errorf("%v\n", err)
		}
	}

	// 录入非法年龄测试大于150
	{
		age, sex := 160, "男"
		bmiResult := 22
		_,err := CalcFatRate(float64(bmiResult), age, sex)
		if err != nil {
			t.Errorf("%v\n",err)
		}
	}

	// 录入非法年龄为负数测试
	{
		age, sex := -30, "男"
		bmiResult := 22
		_,err := CalcFatRate(float64(bmiResult), age, sex)
		if err != nil {
			t.Errorf("%v\n",err)
		}
	}

	// 非男女的性别测试case
	{
		age, sex := -30, "我就不输入女或者男"
		bmiResult := 22
		_,err := CalcFatRate(float64(bmiResult), age, sex)
		if err != nil {
			t.Errorf("%v\n",err)
		}
	}
}

// 3.录入完整的性别、年龄、身高、体重，确保最终获得的健康建议符合预期。
func TestHealthSuggestCase3(t *testing.T) {
	{
		sex := "男"
		age := 27
		heightM := 1.75
		weightKG := 70
		var expectedSuggest string = "偏重"
		var expectedFatRate float64 = 0.17438571428571426

		bmiResult, err := bmi.CalcBMI(float64(weightKG), heightM)
		if err != nil {
			t.Errorf("%v",err)
		}

		bfrActualResult,err:= CalcFatRate(bmiResult,age,sex)

		if err != nil{
			fmt.Errorf("%v",err)
		}

		if bfrActualResult != 0.17438571428571426{
			t.Errorf("FatRate期望结果:%v,实际结果:%v\n",expectedFatRate,bfrActualResult)
		}

		actualSuggest := suggest.Suggest(sex,float64(age),bfrActualResult)
		if actualSuggest != expectedSuggest {
			t.Errorf("实际建议为:%v,期望建议为%v,不符合预期。",actualSuggest,expectedSuggest)
		}
	}
}