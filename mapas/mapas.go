package mapas

import "fmt"

//Clave-Valor
func MostrarMapas() {
	paises := make(map[string]string)
	paises["Espana"] = "Madrid"
	paises["Francia"] = "Paris"

	campeonato := map[string]int{
		"Barcelona": 1,
		"Madrid":    2,
		"Atletico":  3,
	}
	//fmt.Println(campeonato)

	for equipo, puntuacion := range campeonato {
		fmt.Printf("Equipo %s, puntuacion %d \n", equipo, puntuacion)
	}
	delete(campeonato, "Madrid")
	fmt.Println(campeonato)

	//No existe
	puntuaje, existe := campeonato["Bayern"]
	fmt.Printf("El puntuaje capturado es %d, y el equipo es %t \n", puntuaje, existe)
}
