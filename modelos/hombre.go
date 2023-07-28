package modelos

type Hombre struct {
	Persona
}

func (h *Hombre) Respira()       { h.respirando = true }
func (h *Hombre) Comer()         { h.respirando = true }
func (h *Hombre) Pensar()        { h.respirando = true }
func (h *Hombre) Genero() string { return "hombre" }
