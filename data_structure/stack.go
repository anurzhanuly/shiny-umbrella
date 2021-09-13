package datastructure

import (
	"reflect"
)

type Stack struct {
	ItemStack [8]interface{}
	Top       int
}

func (s *Stack) NewStack() *Stack {
	return s
}

func (s *Stack) Push(item interface{}) {
	if len(s.ItemStack) == s.Top {
		s.resize(s.ItemStack[:])
	}
}

func (s *Stack) resize(original []interface{}) []interface{} {

	switch varType := reflect.TypeOf(original).String(); varType {
	case "[]int":
		varCopy := [len(s.ItemStack) * 2]int{}
		varCopy = s.copyValues(original, varCopy)
	case "[]string":
		varCopy := [len(s.ItemStack) * 2]string{}
	}

}

func (s *Stack) copyValues(original, destination []interface{}) []interface{} {
	for index := range original {
		destination[index] = original[index]
	}

	return destination
}
