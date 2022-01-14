package main

type Person struct {
	targetFloor []int
	RequestPosition []int
}

type Persons struct {
	persons []*Person
}

func (p *Person) RequestFloor(resActionFloor int){
	p.RequestPosition = append(p.RequestPosition,resActionFloor)
}

