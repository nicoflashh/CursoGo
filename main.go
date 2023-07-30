package main

import (
	"fmt"

	"github.com/Cursogo/goroutines"
)

func main() {
	canal1 := make(chan bool)
	go goroutines.MiNombreLento("tusmuertos", canal1)
	fmt.Println("Estoy aqui")

	<-canal1

}
