package defer_panic

import (
	// "os"
	"fmt"
	"log"
)

func VemosDefer() { // para ejecutar instrucciones al terminar el programa
	fmt.Println("Este es un primer mensaje")
	defer fmt.Println("Este es el mensaje final")

	fmt.Println("Eeste es el segundo mensaje")
}

func EjemploPanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("ocurrio un error que generó Panic \n %v", reco)
		}
	}()
	a := 1
	if a == 1 {
		panic("Se encontro el valor 1")
	}
}
