package main

import (
	"fmt"
	gobmi "github.com/armstrongli/go-bmi"
	"github.com/spf13/cobra"
)

var (
	Name   string
	Sex    string
	Height float64 //string
	Weight float64 //string
	Age    int
)

func InputFromCobra() {
	cmd := cobra.Command{
		Use:   "healthcheck",                    //如何使用
		Short: "体脂计算器，根据身高，体重，性别，年龄计算体制比，给出建议。", // 短描述
		Long:  "该体脂计算器是基于bmi的体脂计算器进行计算的...",     // 大段的描述
		Run: func(cmd *cobra.Command, args []string) {


			fmt.Println("Name:", Name)
			fmt.Println("Sex:", Sex)
			fmt.Println("Height:", Height)
			fmt.Println("Weight:", Weight)
			fmt.Println("Age:", Age)

			// 计算
			bmi, err := gobmi.BMI(Weight, Height)

			if err != nil {
				fmt.Errorf("bmi结果异常,错误信息:%v\n", err)
			}

			bfr := gobmi.CalcFatRate(bmi, Age, Sex)
			fmt.Printf("bmi:%v,bfr:%v\n", bmi, bfr)
		},
	}

	cmd.Flags().StringVar(&Name, "Name", "", "姓名")
	cmd.Flags().StringVar(&Sex, "Sex", "", "姓别")
	cmd.Flags().Float64Var(&Height, "Height", -1, "身高")
	cmd.Flags().Float64Var(&Weight, "Weight", -1, "体重")
	cmd.Flags().IntVar(&Age, "Age", -1, "年龄")

	// 运行命令行对象
	cmd.Execute()
}

func main() {
	//ArgsExample()
	InputFromCobra()
}
