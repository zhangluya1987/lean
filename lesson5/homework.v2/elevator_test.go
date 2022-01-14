package main

import (
	"testing"
)

func TestCase1(t *testing.T) {
	// 检查楼层总高是否等于期望值
	e := &Elevators{GetElevator(),
	}

	if totalFloor != 5 {
		t.Errorf("期望楼层数是5,但实际是%v,不符合预期。\n", totalFloor)
	} else {
		t.Logf("期望楼层数是5,实际也是%v,符合预期。\n", totalFloor)
	}

	{
		// 检查每层楼电梯的按键状态是否都等于-1,如果是说明没人请求 电梯不动 原地维持等待状态。
		for floor, everyFloorStatus := range e.elevator.floorsElevatorStatus {
			if everyFloorStatus == -1{
				t.Logf("%v楼电梯状态是%v,符合预期(-1表示每层楼都没人请求)。\n", floor+1,everyFloorStatus)
			}else {
				t.Errorf("%v楼电梯是%v,不符合预期(-1表示每层楼都没人请求)。\n", floor+1,everyFloorStatus)
			}
		}
	}
}

func TestCase2(t *testing.T) {
	e := &Elevators{
		elevator: GetElevator(),
	}

	// 获取当前电梯位置
	if e.elevator.defaultStartupPosition != 1{
		t.Errorf("期望当前电梯位置在1层,实际电梯位置在%v层,不符合预期。\n",e.elevator.defaultStartupPosition)
	} else{
		t.Logf("期望当前电梯位置在1层,实际电梯位置也在%v层,符合预期。\n",e.elevator.defaultStartupPosition)
	}

	// 三楼按电梯，电梯向三楼行进，并停在三楼
	{
		person := &Person{
			targetFloor: 0, // 0表示没动作 等待上电梯中
			currOnFloor: 3,
		}
		e.elevator.ToGetTarget(person)
	}
}

func TestCase3(t *testing.T) {

	e := &Elevators{
		elevator: GetElevator(),
	}

	{

		fakePerson1 := &Person{
			targetFloor: 4,
			currOnFloor: 3,
		}

		fakePerson2 := &Person{
			targetFloor: 2,
			currOnFloor: 3,
		}

		//楼层有5层，电梯在3层。上来一些人后，目标楼层:4楼、2楼。电梯先向上到4楼，然后转头到2楼，最后停在2楼。
		e.elevator.ToGetTarget(fakePerson1)
		e.elevator.ToGetTarget(fakePerson2)

	}
}

// 楼层有5层，电梯在3层。上来一些人后，目标楼层： 4楼、5楼、2楼。电梯先向上到4楼，然后到5楼，之后转头到2楼，最后停在2楼。
func TestCase4(t *testing.T) {
	e := &Elevators{
		elevator: GetElevator(),
	}

	{

		fakePerson1 := &Person{
			targetFloor: 4,
			currOnFloor: 3,
		}

		fakePerson2 := &Person{
			targetFloor: 5,
			currOnFloor: 3,
		}
		fakePerson3 := &Person{
			targetFloor: 2,
			currOnFloor: 3,
		}

		//楼层有5层，电梯在3层。上来一些人后，目标楼层:4楼、2楼。电梯先向上到4楼，然后转头到2楼，最后停在2楼。
		e.elevator.ToGetTarget(fakePerson1)
		e.elevator.ToGetTarget(fakePerson2)
		e.elevator.ToGetTarget(fakePerson3)

	}
}















//
//func TestCase3(t *testing.T) {
//	Floors := &Elevators{
//		elevator: GetElevator(),
//	}
//
//	// 1.楼层有5层，电梯在3层。上来一些人后，目标楼层:4楼、2楼。电梯先向上到4楼，然后转头到2楼，最后停在2楼。
//	{
//		// Floors.elevator.currPosition = 3
//		fakePersons := Persons{persons: []Person{
//			{
//				targetFloor: []int{
//					4,
//					2,
//				},
//				currOnFloor: 3,
//			},
//
//		}}
//
//		for _,person := range fakePersons.persons {
//
//			Floors.elevator.Direction(&person)
//		}
//
//
//	}
//
//}