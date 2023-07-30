package main

import (
	"github.com/Cursogo/middleware"
)

func main() {
	/*
		canal1 := make(chan bool)
		go goroutines.MiNombreLento("tusmuertos", canal1)

		defer func() {
			<-canal1
		}()
		fmt.Println("Estoy aqui")
	*/

	middleware.MiMiddleware()
}
