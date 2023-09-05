package LinkedList

import (
	"errors"
	"fmt"
)

type List[tipo any] interface {
	Add(value tipo)             //
	AddPos(value tipo, pos int) //
	Update(value tipo, pos int) //
	RemoveLast()                //
	Remove(pos int)
	Get(pos int) tipo
	Size() int //
}

type Node[tipo any] struct {
	value tipo
	next  *Node[tipo]
}

type LinkedList[tipo any] struct {
	head     *Node[tipo]
	inserted int
}

func (list *LinkedList[tipo]) Add(val tipo) {
	node := Node[tipo]{val, nil}
	if list.inserted == 0 {
		list.head = &node
		list.inserted++
		return
	}
	last := list.head
	for ; last.next != nil; last = last.next {
	}
	last.next = &node
	list.inserted++
}

func (list *LinkedList[tipo]) AddPos(val tipo, pos int) error {
	if pos < 0 || pos > list.inserted {
		return errors.New("Índice fora de alcance")
	}
	node := Node[tipo]{val, nil}
	last := list.head
	if pos == 0 {
		node.next = last
		list.head = &node
		list.inserted++
		return nil
	}
	for i := 1; i < pos; i++ {
		last = last.next
	}
	node.next = last.next
	last.next = &node
	list.inserted++
	return nil
}

func (list *LinkedList[tipo]) Update(val tipo, pos int) {
	if pos >= list.inserted {
		panic("Índice fora de alcance!")
	}
	nodePtr := list.head
	for i := 0; i < pos; i++ {
		nodePtr = nodePtr.next
	}
	nodePtr.value = val
}

func (list *LinkedList[tipo]) RemoveLast() {
	if list.inserted == 0 {
		panic("Não há elemento a se remover!")
	}
	if list.inserted == 1 {
		list.head = nil
		list.inserted--
		return
	}
	nodePtr := list.head
	for ; nodePtr.next.next != nil; nodePtr = nodePtr.next {
	}
	nodePtr.next = nil
	list.inserted--
}

func (list *LinkedList[tipo]) Remove(pos int) {
	if pos < 0 || pos >= list.inserted {
		panic("Índice fora de alcance!")
	}
	if pos == 0 {
		list.head = list.head.next
		list.inserted--
		return
	}
	nodePtr := list.head
	for i := 0; i < pos-1; i++ {
		nodePtr = nodePtr.next
	}
	nodePtr.next = nodePtr.next.next
	list.inserted--
}

func (list *LinkedList[tipo]) Get(pos int) tipo {
	if pos < 0 || pos >= list.inserted {
		panic("Índice fora de alcance!")
	}
	nodePtr := list.head
	for i := 0; i < pos; i++ {
		nodePtr = nodePtr.next
	}
	return nodePtr.value
}

func (list *LinkedList[tipo]) Size() int {
	return list.inserted
}

func (list *LinkedList[tipo]) Display() {
	out := "{"
	nodePtr := list.head
	for ; nodePtr != nil; nodePtr = nodePtr.next {
		if nodePtr.next != nil {
			out = fmt.Sprint(out, nodePtr.value, ",")
		} else {
			out = fmt.Sprint(out, nodePtr.value, "}")
		}
	}
	fmt.Println(out)
}

func Demo() {
	var teste LinkedList[int]
	fmt.Println("Teste inicialização")
	fmt.Println(teste)
	teste.Display()
	teste.Add(3)
	teste.Add(2)
	teste.Add(1)
	fmt.Println("Teste Add [3 .. 1]")
	teste.Display()
	fmt.Println(teste)
	teste.AddPos(0, 3)
	fmt.Println("Teste AddPos (0,3), (5,0), (9,3)")
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
	teste.Remove(2)
	fmt.Println("Teste Remove 2, 4, 0")
	teste.Display()
	fmt.Println(teste)
	teste.Remove(4)
	teste.Display()
	fmt.Println(teste)
	teste.Remove(0)
	teste.Display()
	fmt.Println(teste)
	teste.RemoveLast()
	fmt.Println("Teste RemoveLast")
	teste.Display()
	fmt.Println(teste)
	fmt.Println(teste.Get(1))
	fmt.Println("Teste Get 1")
	teste.Display()
	fmt.Println(teste)
	fmt.Println("Teste Size")
	fmt.Println(teste.Size())
}
