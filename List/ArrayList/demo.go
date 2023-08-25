package ArrayList

import (
	"fmt"
)

func Demo() {

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
