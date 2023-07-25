package variables

import (
	"fmt"
)

func MostrarVariables() {
	var intcomun int
	intde32 := int32(10)
	intde64 := int64(100)

	fmt.Println("Intcomun: ", intcomun)
	fmt.Println("Int32: ", intde32)
	fmt.Println("Int64: ", intde64)

}
