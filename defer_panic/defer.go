package defer_panic

import (
	"fmt"
	"log"
)

func VamosDefer() {
	fmt.Println("Primer mensaje")
	defer fmt.Println("Mensaje final")
	fmt.Println("Segundo mensaje")

}

func EjemploPanic() {

	defer func() {
		reco := recover()
		log.Fatalf("ocurrio un error que genero Panic:  %v", reco)
	}()
	a := 1
	if a == 1 {
		panic("Se encontro valor 1")
	}
}
