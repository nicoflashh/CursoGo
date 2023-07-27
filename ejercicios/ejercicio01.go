package ejercicios

import (
	"strconv"
)

func ConvertirString(numero string) (int, string) {
	numeroint, err := strconv.Atoi(numero)
	if err != nil {
		return 0, "Hubo error" + err.Error()
	}
	if numeroint > 100 {
		return numeroint, "Es mayor que 100"
	} else {
		return numeroint, "Es menor que 100"
	}

}
