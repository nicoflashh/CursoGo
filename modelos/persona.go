package modelos

type Persona struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	vivo       bool
}

func (p *Persona) Respira()       { p.respirando = true }
func (p *Persona) Comer()         { p.respirando = true }
func (p *Persona) Pensar()        { p.respirando = true }
func (p *Persona) Genero() string { return "mujer" }
