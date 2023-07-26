package main

import (
	"fmt"

	"github.com/Cursogo/variables"
)

func main() {
	estado, texto := variables.ConviertoaTexto(293838)
	fmt.Println(estado)
	fmt.Println(texto)
}
