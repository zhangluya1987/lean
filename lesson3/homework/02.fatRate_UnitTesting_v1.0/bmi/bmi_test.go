package bmi

import (
	"testing"
)

// TestCase1.录入正常身高、体重，确保计算结果符合预期；
func TestBMICase1(t *testing.T) {
	// 实际计算得出 体重70kg,身高1.75m bmi计算结果为 22.857142857142858
	weight, height := 70.0, 1.75
	bmi, err := CalcBMI(weight, height)

	if err != nil {
		t.Errorf("您输入的身高%v,或者体重%v 异常,请重新输入", height, weight)
	}

	if bmi != 22.857142857142858 {
		t.Errorf("预期结果是22.9,但计算结果是%v,不符合预期。\n", bmi)
	}

}

// TestBMICase2.录入0或负数身高，返回错误；
func TestBMICase2(t *testing.T) {
	// 录入0身高,返回错误。
	{
		_, err := CalcBMI(70, 0)
		if err != nil {
			t.Errorf("身高数值输入异常: %v\n", err)
		}
	}

	//录入负数身高，返回错误。
	{
		_, err := CalcBMI(70, -2)
		if err != nil {
			t.Errorf("身高数值输入异常: %v\n", err)
		}
	}
}

// TestBMICase3.录入0或负数体重，返回错误。
func TestBMICase3(t *testing.T){
	// 录入0体重，返回错误。
	{
		_, err := CalcBMI(0, 1.75)
		if err != nil {
			t.Errorf("体重数值输入异常: %v\n", err)
		}
	}

	//录入负数体重，返回错误。
	{
		_, err := CalcBMI(-2, 1.75)
		if err != nil {
			t.Errorf("体重数值输入异常: %v\n", err)
		}
	}
}
