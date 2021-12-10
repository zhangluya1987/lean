package input

import "fmt"

var (
	X1, Y1 float64
	X2, Y2 float64
	X3, Y3 float64
	X4, Y4 float64
)

func Input() {
	fmt.Println("请输入(空格隔开):X1 Y1")
	fmt.Scanln(&X1, &Y1)
	fmt.Println("请输入X2 Y2")
	fmt.Scanln(&X2, &Y2)
	fmt.Println("请输入X3 Y3")
	fmt.Scanln(&X3, &Y3)
	fmt.Println("请输入X4 Y4")
	fmt.Scanln(&X4, &Y4)
}
