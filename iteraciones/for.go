package iteraciones

import (
	"fmt"
)

func Iterar() {
	for i := 0; i < 10; i += 1 { // i += 5  para dar saltos de 5
		if i == 6 {
			continue // vuelve arriba y omite las lineas inferiores
		}
		fmt.Println(i)
	}

}
