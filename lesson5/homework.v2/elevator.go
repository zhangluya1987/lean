package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

const totalFloor = 5

type Elevator struct {
	defaultStartupPosition int
	currPosition           int
	tagetFloor             int
	lastPersonResFloor     int
	defaultStatus          int
	floorsElevatorStatus   []int
}

type Elevators struct {
	elevator *Elevator
}

func GetElevator() *Elevator {
	return &Elevator{
		defaultStartupPosition: 1,
		currPosition:           1,
		tagetFloor:             -1,
		lastPersonResFloor:     -1,
		defaultStatus:          -1,
		floorsElevatorStatus: []int{
			-1,-1,-1,-1,-1,
		},
	}
}

func (e *Elevator) ToGetTarget(person *Person) {
	e.tagetFloor = person.currOnFloor
	fmt.Printf("我当前在%v层,我要去%v层接人.\n", e.currPosition, e.tagetFloor)

	//耗费时长
	getCostTime := person.currOnFloor - e.currPosition
	costTime := int(math.Abs(float64(getCostTime)))
	fmt.Printf("上楼花费%vs\n", costTime)
	time.Sleep(time.Second * time.Duration(costTime))

	fmt.Printf("已到达%v层\n",e.tagetFloor)
	if person.targetFloor == 0 {
		fmt.Printf("等待接人中,电梯暂停在%v层.\n",e.tagetFloor)
	}else {
		e.SendToTargetFloor(person)
	}
}

func (e *Elevator) SendToTargetFloor(person *Person) {
	e.currPosition = person.currOnFloor
	e.tagetFloor = person.targetFloor
	fmt.Printf("我当前在%v层,我要送人去%v层.\n", e.currPosition, e.tagetFloor)

	//耗费时长
	getCostTime := person.targetFloor - e.currPosition
	costTime := int(math.Abs(float64(getCostTime)))
	time.Sleep(time.Second * time.Duration(costTime))
	fmt.Printf("送人到目的楼层花费%vs\n", costTime)
	fmt.Printf("已送人到%v层\n",person.targetFloor)
	fmt.Println(strings.Repeat("-",50))
	e.currPosition = person.targetFloor
}
