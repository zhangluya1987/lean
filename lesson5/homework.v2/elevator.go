package main

import (
	"fmt"
	"strings"
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
		lastPersonResFloor:     1,
		defaultStatus:          -1,
		floorsElevatorStatus: []int{
			-1,-1,-1,-1,-1,
		},
	}
}


func (e *Elevator) ToGetTarget(person *Person) {

	if e.currPosition == 1{
		e.currPosition = person.currOnFloor
		fmt.Printf("我要去%v层接人\n",person.currOnFloor)
	}

	if person.targetFloor == 0 {
		fmt.Printf("我当前在%v层,电梯等待接人中...\n", e.currPosition)
	}else {
		fmt.Printf("我当前在%v层,去%v层.\n", e.currPosition, person.targetFloor)
	}

	e.currPosition = person.targetFloor
	if person.targetFloor == 0 {
		fmt.Printf("等待接人中,电梯暂停...\n")
	}else {
		e.SendToTargetFloor(person)
	}
}

func (e *Elevator) SendToTargetFloor(person *Person) {
	fmt.Printf("我要送人去%v层.\n",person.targetFloor)
	e.lastPersonResFloor = person.targetFloor

	fmt.Printf("已送人到%v层\n",person.targetFloor)
	fmt.Println(strings.Repeat("-",50))
	e.currPosition = person.targetFloor
}
