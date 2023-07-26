package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string
var State bool
var Sueldo float32
var Fecha time.Time

func RestoVariables() {
	Nombre = "Pedro"
	State = true
	Sueldo = 2900.9
	Fecha = time.Now()
	fmt.Println(Nombre)
	fmt.Println(State)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)
}

func ConviertoaTexto(numero int) (bool, string) {
	var texto string
	texto = strconv.Itoa(numero)
	return true, texto
}
