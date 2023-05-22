package middleware

import "fmt"

func sumar(a, b int) int {
	return a + b
}
func restar(a, b int) int {
	return a - b
}
func multiplicar(a, b int) int {
	return a * b
}

func MiMiddleware() {
	fmt.Println("Inicio")
	result := operacionesMidd(sumar)(2, 3) //la función que  retorna operacionesMidd (en este caso sumar), la ejecuta y le pasa los parametros, en este caso, 2 y 3
	fmt.Println(result)
	result = operacionesMidd(restar)(21, 5)
	fmt.Println(result)
	result = operacionesMidd(multiplicar)(5, 6)
	fmt.Println(result)
}

// función que recibe como parametro una función y retorna una función
func operacionesMidd(funcion func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("inicio de Operación")
		return funcion(a, b)
	}
}
