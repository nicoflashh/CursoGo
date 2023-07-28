package ejerciciointerfaces

import (
	"fmt"

	"github.com/Cursogo/interfaces"
)

func PersonaRespira(p1 interfaces.Persona) {
	p1.Respira()
	fmt.Printf("Genero: %s, y respiro", p1.Genero())
}
