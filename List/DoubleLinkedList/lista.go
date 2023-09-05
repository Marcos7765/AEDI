package main

import (
	"errors"
	"fmt"
)

type List[tipo any] interface {
	Add(value tipo)             //
	AddPos(value tipo, pos int) //
	Update(value tipo, pos int)
	RemoveLast()
	Remove(pos int)
	Get(pos int) tipo
	Size() int //
}

type Node[tipo any] struct {
	value tipo
	next  *Node[tipo]
	prev  *Node[tipo]
}

type DoubleLinkedList[tipo any] struct {
	head     *Node[tipo]
	inserted int
	tail     *Node[tipo]
}

func (list *DoubleLinkedList[tipo]) Add(val tipo) {
	node := Node[tipo]{val, nil, list.tail}
	if list.inserted == 0 {
		list.head = &node
		list.tail = &node
		list.inserted++
		return
	}
	list.tail.next = &node
	list.tail = &node
	list.inserted++
}

func (list *DoubleLinkedList[tipo]) AddPos(val tipo, pos int) error {
	if pos < 0 || pos > list.inserted {
		return errors.New("Índice fora de alcance")
	}
	node := Node[tipo]{val, nil, nil}
	if pos == 0 {
		node.next = list.head
		list.head.prev = &node
		list.head = &node
		list.inserted++
		return nil
	}
	if pos == list.inserted {
		node.prev = list.tail
		list.tail.next = &node
		list.tail = &node
		list.inserted++
		return nil
	}
	var last *Node[tipo]
	if pos > list.inserted/2 {
		last = list.tail
		for i := list.inserted - 1; i > pos; i-- {
			last = last.prev
		}
		node.next = last
		node.prev = last.prev
		last.prev.next = &node
		last.prev = &node

	} else {
		last = list.head
		for i := 1; i < pos; i++ {
			last = last.next
		}
		node.next = last.next
		node.prev = last
		last.next.prev = &node
		last.next = &node
	}
	list.inserted++
	return nil
}

func (list *DoubleLinkedList[tipo]) Update(val tipo, pos int) {
	if pos >= list.inserted {
		panic("Índice fora de alcance!")
	}
	var last *Node[tipo]
	if pos > list.inserted/2 {
		last = list.tail
		for i := list.inserted - 1; i > pos; i-- {
			last = last.next
		}
	} else {
		last = list.head
		for i := 0; i < pos; i++ {
			last = last.next
		}
	}
	last.value = val
}

func (list *DoubleLinkedList[tipo]) RemoveLast() {
	if list.inserted == 0 {
		panic("Não há elemento a se remover!")
	}
	if list.inserted == 1 {
		list.head = nil
		list.tail = nil
		list.inserted--
		return
	}
	list.tail = list.tail.prev
	list.tail.next = nil
	list.inserted--
}

func (list *DoubleLinkedList[tipo]) Remove(pos int) {
	if pos < 0 || pos >= list.inserted {
		panic("Índice fora de alcance!")
	}
	if list.inserted == 1 {
		list.head = nil
		list.tail = nil
		list.inserted--
		return
	}
	if pos == 0 {
		list.head = list.head.next
		list.inserted--
		return
	}
	if pos == list.inserted-1 {
		list.tail = list.tail.prev
		list.tail.next = nil
		list.inserted--
		return
	}

	var last *Node[tipo]
	if pos >= list.inserted/2 {
		last = list.tail
		for i := list.inserted - 2; i > pos; i-- {
			last = last.prev
		}
		last.prev.prev.next = last
		last.prev = last.prev.prev
	} else {
		last = list.head
		for i := 1; i < pos; i++ {
			last = last.next
		}
		last.next.next.prev = last
		last.next = last.next.next
	}
	list.inserted--
}

func (list *DoubleLinkedList[tipo]) Get(pos int) tipo {
	if pos < 0 || pos >= list.inserted {
		panic("Índice fora de alcance!")
	}
	var last *Node[tipo]
	if pos > list.inserted/2 {
		last = list.tail
		for i := list.inserted - 1; i > pos; i-- {
			last = last.next
		}
	} else {
		last = list.head
		for i := 1; i < pos; i++ {
			last = last.next
		}
	}

	return last.value
}

func (list *DoubleLinkedList[tipo]) Size() int {
	return list.inserted
}

func (list *DoubleLinkedList[tipo]) Display() {
	if list.inserted == 0 {
		fmt.Println("head = ", list.head, ", tail = ", list.tail)
		return
	}
	out := "{"
	nodePtr := list.head
	for ; nodePtr != nil; nodePtr = nodePtr.next {
		if nodePtr.next != nil {
			out = fmt.Sprint(out, nodePtr.value, ",")
		} else {
			out = fmt.Sprint(out, nodePtr.value, "}")
		}
	}
	out = fmt.Sprint(out, " tail = ", list.tail.value)
	fmt.Println(out)
}

func main() {
	var teste DoubleLinkedList[int]
	fmt.Println("Teste inicialização")
	fmt.Println(teste)
	teste.Display()
	teste.Add(3)
	teste.Add(2)
	teste.Add(1)
	fmt.Println("Teste Add [3 .. 1]")
	teste.Display()
	fmt.Println(teste)
	fmt.Println("Teste AddPos (0,3), (5,0), (9,3)")
	teste.AddPos(0, 3)
	teste.Display()
	fmt.Println(teste)
	teste.AddPos(5, 0)
	teste.Display()
	fmt.Println(teste)
	teste.AddPos(9, 3)
	teste.Display()
	fmt.Println(teste)
	fmt.Println("Teste Update (7,3)")
	teste.Update(7, 3)
	teste.Display()
	fmt.Println(teste)
	fmt.Println("Teste Remove 2, 4, 0")
	teste.Remove(2)
	teste.Display()
	fmt.Println(teste)
	teste.Remove(4)
	teste.Display()
	fmt.Println(teste)
	teste.Remove(0)
	teste.Display()
	fmt.Println("Teste RemoveLast")
	fmt.Println(teste)
	teste.RemoveLast()
	fmt.Println("Teste Get 1")
	teste.Display()
	fmt.Println(teste)
	fmt.Println(teste.Get(1))
	fmt.Println("Teste Size")
	teste.Display()
	fmt.Println(teste)
	fmt.Println(teste.Size())
}
