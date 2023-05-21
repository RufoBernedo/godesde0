package modelos

type Mujer struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	vivo       bool
}

func (mujer *Mujer) Respirar()      { mujer.respirando = true }
func (mujer *Mujer) Comer()         { mujer.comiendo = true }
func (mujer *Mujer) Pensar()        { mujer.pensando = true }
func (mujer *Mujer) Sexo() string   { return "Mujer" }
func (mujer *Mujer) EstaVivo() bool { return mujer.vivo }
