package arreglosslice

import "fmt"

var tablas []int = []int{2, 5, 4}
var arreglo [10]int = [10]int{6, 1, 4, 7, 8, 3, 4, 5}

func MuestroSlice() {
	fmt.Println(tablas)
	porcion := arreglo[3:]
	porcion2 := arreglo[:5]
	porcion3 := arreglo[6:7]
	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func Capacidad() {
	elementos := make([]int, 5, 20)

	fmt.Printf("Largo %d Capacidad %d", len(elementos), cap(elementos))

	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}

	fmt.Printf("Largo %d Capacidad %d", len(nums), cap(nums))
}
