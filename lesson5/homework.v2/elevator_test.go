package main

import (
	"testing"
)

func TestCase1(t *testing.T) {

	e := &Elevators{elevators: GetElevator()}
	if len(e.elevators.totalFloors) != 5{
		t.Logf("总楼层期望值是5,但实际是%v,不符合预期.\n",len(e.elevators.totalFloors))
	} else {
		t.Logf("总楼层期望值是5,实际也是%v,符合预期.\n",len(e.elevators.totalFloors))
	}

	for index,floor := range e.elevators.totalFloors{
		if floor[0] != -1 {
			t.Logf("第%v层电梯期望的状态是-1,实际是%v,不符合预期\n",index+1,floor[0])
		}else {
			t.Logf("第%v的电梯期望状态是-1,实际也是%v,符合预期\n",index+1,floor[0])
		}
	}
}

func TestCase2(t *testing.T){
	e := &Elevators{elevators: GetElevator()}
	person := &Person{
		targetFloor:     []int{},
		RequestPosition: []int{},
	}

	fakeInputRes := 3
	person.RequestFloor(fakeInputRes)
	e.elevators.resQueue =append(e.elevators.resQueue,fakeInputRes)
	e.elevators.GoTo(person)
	if len(e.elevators.resQueue) ==1 && len(e.elevators.sendQueue) == 0 {
		t.Logf("电梯停靠在%v层,按目标楼层键的人数为%v,电梯暂停中。",e.elevators.resQueue[0],len(e.elevators.sendQueue))
	}
}

func TestCase3(t *testing.T) {
	{
		e := &Elevators{elevators: GetElevator()}
		person := &Person{
			targetFloor:     []int{},
			RequestPosition: []int{},
		}

		person.RequestFloor(4)
		person.RequestFloor(2)
		e.elevators.LastPos(3)
		e.elevators.sendQueue = append(e.elevators.sendQueue, person.RequestPosition...)
		e.elevators.ChangeStatus(*person)
		e.elevators.Schedule(person)
		e.elevators.GoTo(person)
		//todo 单测部分待补充
	}
}

func TestCase4(t *testing.T) {
	e := &Elevators{elevators: GetElevator()}
	person := &Person{
		targetFloor:     []int{},
		RequestPosition: []int{},
	}

	{
		person.RequestFloor(4)
		person.RequestFloor(5)
		person.RequestFloor(2)
		e.elevators.LastPos(3)
		e.elevators.sendQueue = append(e.elevators.sendQueue, person.RequestPosition...)
		e.elevators.ChangeStatus(*person)
		e.elevators.Schedule(person)
		e.elevators.GoTo(person)
		//todo 单测部分待补充
	}
}