package LinkedStack

import (
	"errors"
	"fmt"
)

type Stack[tipo any] interface {
	Push(value tipo)
	Pop()
	Get(pos int) tipo
	Size() int
}

type Node[tipo any] struct {
	value tipo
	next  *Node[tipo]
}

type LinkedStack[tipo any] struct {
	top     *Node[tipo]
	size int
}

func (stack *LinkedStack[tipo]) Push(val tipo) {
	node := Node[tipo]{val, stack.top}
	stack.top = &node
	stack.size++
}

func (stack *LinkedStack[tipo]) Pop() tipo {
	if stack.size == 0 {
		panic("Não há elemento a se remover!")
	}
	res := stack.top.value
	stack.top = stack.top.next
	stack.size--
	return res
}

func (stack *LinkedStack[tipo]) Top() tipo {
	if stack.size == 0 {
		panic("A stack está vázia")
	}
	return stack.top.value
}

func (stack *LinkedStack[tipo]) Size() int {
	return stack.size
}

func (stack *LinkedStack[tipo]) Empty() bool {
	return stack.size == 0
}

func teste(){
	var teste LinkedStack[string]
	fmt.prin
}