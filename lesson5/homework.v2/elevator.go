package main

import (
	"fmt"
)

var(
	up []int
	down []int
)

type Elevator struct {
	totalFloors [][]int //一维 楼层 2维状态
	sendQueue []int
	firstStartUpFloor int
	lastPosition int
	resQueue []int
}

type Elevators struct {
	elevators *Elevator
}

func GetElevator() *Elevator {
	return &Elevator{
		totalFloors:       [][]int{
			{
				-1,
			},
			{
				-1,
			},
			{
				-1,
			},
			{
				-1,
			},
			{
				-1,
			},
		},
		sendQueue:             []int{},
		firstStartUpFloor: 1,
		lastPosition: 1,
		resQueue: []int{},
	}
}

func (e *Elevator) LastPos(res int){
	e.lastPosition = res
}

func (e *Elevator) Tag(res int){
	fmt.Printf("送人去%v层\n",res)
}
func (e *Elevator) Pause(){
	fmt.Println("电梯暂停.")
}

func (e *Elevator) OpenDoor(){
	fmt.Println("电梯开门.")
}

func (e *Elevator) CloseDoor(){
	fmt.Println("电梯关门.")
	if len(e.resQueue) == 0 {
		e.Pause()
	}
}

func (e *Elevator) NextFloor(){
	fmt.Println("继续原方向前进.")
}

func (e *Elevator) GoTo(person *Person){
	if len(e.resQueue)==1{
		fmt.Printf("我要去%v层.\n",e.resQueue[0])
		e.OpenDoor()
		if len(e.sendQueue)==0{
			fmt.Println("人进入电梯啦,但是没有人按目标楼层按键。")
		}
		e.CloseDoor()
		e.Pause()
		e.lastPosition = e.resQueue[0]

	}

	if len(e.sendQueue) >1 {
		if len(up) >= len(down) {
			for _, v := range up {
				e.Tag(v)
				e.OpenDoor()
				e.CloseDoor()
				e.NextFloor()
				e.lastPosition = v
				fmt.Printf("当前在%v层\n", e.lastPosition)
			}

			for _, v := range down {
				e.Tag(v)
				e.OpenDoor()
				e.CloseDoor()
				e.NextFloor()
				e.lastPosition = v
				fmt.Printf("当前在%v层\n", e.lastPosition)
			}
		}else{
			for _,v := range down{
				e.Tag(v)
				e.OpenDoor()
				e.CloseDoor()
				e.NextFloor()
				e.lastPosition = v
				fmt.Printf("当前在%v层\n",e.lastPosition)
			}

			for _,v := range up{
				e.Tag(v)
				e.OpenDoor()
				e.CloseDoor()
				e.NextFloor()
				e.lastPosition = v
				fmt.Printf("当前在%v层\n",e.lastPosition)
			}
			}

		}
	}

func (e *Elevator) ChangeStatus(person Person) {
	for _,value := range e.sendQueue{
		e.totalFloors[value-1][0]=value
	}
}

func (e *Elevator) Schedule(person *Person) {
	{
		for _, target := range e.sendQueue {
			if target > e.lastPosition {
				up = append(up, target)
			} else {
				down = append(down, target)
			}
		}
	}

}