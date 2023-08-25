//go:build direto

package lista

import (
	"errors"
	"fmt"
)

type List[tipo any] interface {
	Add(value tipo)             //
	AddPos(value tipo, pos int) //
	Update(value tipo, pos int) //
	RemoveLast()                //
	Remove(pos int)             //
	Get(pos int) tipo
	Size() int //
}

type ArrayList[tipo any] struct {
	values   []tipo
	inserted int
}

func (list *ArrayList[tipo]) Add(val tipo) {
	defer func() {
		if r := recover(); r != nil {
			size := 2 * list.inserted
			if size == 0 {
				size = 10
			}
			var placeholder []tipo = make([]tipo, size)
			copy(placeholder, list.values)
			list.values = placeholder
			list.Add(val)
		}
	}()
	list.values[list.inserted] = val
	list.inserted++
}

func (list *ArrayList[tipo]) AddPos(val tipo, pos int) error {
	defer func() {
		if r := recover(); r != nil {
			size := 2 * list.inserted
			if size == 0 {
				size = 10
			}
			var placeholder []tipo = make([]tipo, size)
			copy(placeholder, list.values)
			list.values = placeholder
			list.AddPos(val, pos)
		}
	}()
	if pos < 0 || pos > list.inserted {
		return errors.New("Índice fora de alcance")
	}
	valph1 := list.values[pos]
	valph2 := list.values[list.inserted] //forçar o realloc antes de mudar valores
	for i := pos; i < list.inserted; i++ {
		valph2 = list.values[i+1]
		list.values[i+1] = valph1
		valph1 = valph2
	}
	list.values[pos] = val
	list.inserted++
	return nil
}

func (list *ArrayList[tipo]) Update(val tipo, pos int) {
	if pos >= list.inserted {
		panic("Índice fora de alcance!")
	}
	list.values[pos] = val
}

func (list *ArrayList[tipo]) RemoveLast() {
	if list.inserted == 0 {
		panic("Não há elemento a se remover!")
	}
	var placeholder tipo
	list.values[list.inserted-1] = placeholder
	list.inserted--
}

func (list *ArrayList[tipo]) Remove(pos int) {
	if pos < 0 || pos >= list.inserted {
		panic("Índice fora de alcance!")
	}
	for i := pos; i < list.inserted-1; i++ {
		list.values[i] = list.values[i+1]
	}
	var placeholder tipo
	list.inserted--
	list.values[list.inserted] = placeholder
}

func (list *ArrayList[tipo]) Get(pos int) tipo {
	if pos < 0 || pos >= list.inserted {
		panic("Índice fora de alcance!")
	}
	return list.values[pos]
}

func (list *ArrayList[tipo]) Size() int {
	return list.inserted
}

func main() {

	var teste ArrayList[int]
	fmt.Println("ArrayList inicializada:")
	fmt.Println("\t", teste)

	for i := 1; i < 6; i++ {
		teste.Add(i * 11)
	}
	fmt.Println("ArrayList após chamar Add para 5 múltiplos de 11:")
	fmt.Println("\t", teste)

	teste.AddPos(9, 2)
	fmt.Println("ArrayList após chamar AddPos(9, 2)")
	fmt.Println("\t", teste)

	teste.Update(7777, 4)
	fmt.Println("ArrayList após chamar Update(7777, 4)")
	fmt.Println("\t", teste)

	teste.RemoveLast()
	fmt.Println("ArrayList após chamar RemoveLast()")
	fmt.Println("\t", teste)

	teste.Remove(0)
	fmt.Println("ArrayList após chamar Remove(0)")
	fmt.Println("\t", teste)

	fmt.Println("Valor obtido por Get(3):\n\t", teste.Get(3))

	fmt.Println("Valor obtido por Size():\n\t", teste.Size())
}
