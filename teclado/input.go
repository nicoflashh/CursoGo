package teclas

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var numero2 int
var texto string
var err error

func IngresarNumeros() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese numero 1: ")
	//Revisamos si el usuario ha introducido algo
	if scanner.Scan() {
		numero1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			//Abortar aplicacion
			panic("El dato ingresado es incorrecto" + err.Error())
		}
	}

	fmt.Println("Ingrese numero 2: ")
	//Revisamos si el usuario ha introducido algo
	if scanner.Scan() {
		numero2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			//Abortar aplicacion
			panic("El dato ingresado es incorrecto" + err.Error())
		}
	}

	fmt.Println("Ingrese texto: ")
	//Revisamos si el usuario ha introducido algo
	if scanner.Scan() {
		texto = scanner.Text()

	}
	fmt.Println(texto, numero1*numero2)
}
