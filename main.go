package main

import (
	"fmt"

	"github.com/Cursogo/goroutines"
)

func main() {

	go goroutines.MiNombreLento("tusmuertos")

	fmt.Println("Estoy aqui")
	var x string
	fmt.Scanln(&x)
}
