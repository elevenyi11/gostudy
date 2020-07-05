package main

import (
	"reflect"

	"github.com/pkg/errors"
)

type Stack struct {
	Values    []interface{}
	ValueType reflect.Type
}

func NewStack(valueType reflect.Type) *Stack {
	return &Stack{Values: make([]interface{}, 0), ValueType: valueType}
}

func (Stack *Stack) isAcceptableValue(value interface{}) bool {
	if value == nil || reflect.TypeOf(value) != Stack.ValueType {
		return false
	}
	return true
}

func (stack *Stack) Push(v interface{}) bool {
	if !stack.isAcceptableValue(v) {
		return false
	}
	stack.Values = append(stack.Values, v)
	return true
}

func (stack *Stack) Pop() (interface{}, error) {
	if stack == nil || len(stack.Values) == 0 {
		return nil, errors.New("stack empty")
	}

	v := stack.Values[len(stack.Values)-1]
	stack.Values = stack.Values[:len(stack.Values)-1]
	return v, nil
}

func (stack *Stack) Top() (interface{}, error) {
	if stack == nil || len(stack.Values) == 0 {
		return nil, errors.New("stack empty")
	}
	return stack.Values[len(stack.Values)-1], nil
}

func (stack *Stack) Len() int {
	return len(stack.Values)
}

func (stack *Stack) Empty() bool {
	if stack == nil || len(stack.Values) == 0 {
		return true
	}
	return false
}

func (stack *Stack) StackValueType() reflect.Type {
	return stack.ValueType
}
