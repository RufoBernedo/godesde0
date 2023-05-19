package funciones

import "fmt"

func Exponencia(valor int) {
	fmt.Println(valor)
	if valor > 100000000 {
		return
	}
	Exponencia(valor * 2)
}
