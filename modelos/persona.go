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
func (p *Persona) Comer()         { p.comiendo = true }
func (p *Persona) Pensar()        { p.pensando = true }
func (p *Persona) Genero() string { return "mujer" }
func (p *Persona) EstaVivo() bool { return true }
