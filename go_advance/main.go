package main

import _case "go_advance/case"

func main() {
	t := &_case.Teacher{}
	s := &_case.Student{}
	_case.PersonCase(t)
	_case.PersonCase(s)
	_case.PersonCase1(t)
	_case.TryCatchCase()
	_case.DeferCase()
	_case.DeferCase1()
}
