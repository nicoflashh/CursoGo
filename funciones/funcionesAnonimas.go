package funciones

import "fmt"

func Calculos() {
	var numero3 int = 50
	var numero4 int = 60
	calculo := func(numero1 int, numero2 int) int {
		return numero1 + numero2 + numero3 + numero4
	}
	fmt.Println(calculo(10, 15))

	calculo = func(numero1 int, numero2 int) int {
		return numero1 * numero2 * numero3
	}

	fmt.Println(calculo(10, 15))
}
