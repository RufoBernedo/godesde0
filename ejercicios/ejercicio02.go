package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func CrearTablaNumero() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Ingrese un número: ")
		if scanner.Scan() {
			numero, err := strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			}
			for i := 1; i < 11; i++ {
				fmt.Printf("%d x %d = %d \n", numero, i, i*numero)
			}
			break
		}
	}
}
