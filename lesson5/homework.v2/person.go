package main

import "fmt"

type Person struct {
	targetFloor int
	currOnFloor int
}

type Persons struct {
	persons []Person
}

func (p *Person) Action(person *Person) {
	p.targetFloor = person.targetFloor
	p.currOnFloor = person.currOnFloor
	fmt.Printf("我当前在%v层,我要去%v层.\n",p.currOnFloor,p.targetFloor)
}