package ejer_interfaces

import (
	"fmt"
	"godesde0/interfaces"
)

func HumanosRespirando(humano interfaces.Humano) {
	humano.Respirar()
	fmt.Printf("Soy un/a %s, y estoy respirando \n", humano.Sexo())
}
