//go:build direto

package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 1; i <= 8; i++ {
		var recorde int64 = -1
		for j := 1; j <= 20; j++ {
			placeholder := testar(int(math.Pow(10, float64(i))))
			if recorde == -1 || recorde > placeholder {
				recorde = placeholder
			}
		}
		fmt.Println("Recorde para ", math.Pow(10, float64(i)), " elementos, "+
			"direto: ", recorde)
	}
}
