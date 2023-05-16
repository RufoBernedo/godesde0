package ejercicios

import (
	"strconv"
)

func ConvierteaInt(texto string) (int, string) {

	numero, error := strconv.Atoi(texto)

	if error != nil {
		// ... handle error
		// panic(error)
		return 0, "Hubo un error " + error.Error()
	}
	if numero > 100 {
		return numero, "Es mayor a 100"
	} else {
		return numero, "Es menor a 100"
	}
}
