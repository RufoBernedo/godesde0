package main

// import "godesde0/iteraciones"
import (
	// "godesde0/files"
	// "godesde0/funciones"
	// "godesde0/arreglos_slices"
	// "godesde0/mapas"
	// "godesde0/users"
	// e "godesde0/ejer_interfaces"
	// modelos "godesde0/modelos"
	// dp "godesde0/defer_panic"
	"fmt"
	"godesde0/goroutines"
)

func main() {
	// variables.MuestroEnteros()
	// variables.RestoVariables()

	// estado, texto := variables.ConvierteaTexto(1234)

	// fmt.Println(estado)
	// fmt.Println(texto)

	/*if os := runtime.GOOS; os == "Linux." || os == "OS X." || os == "darwin" {
		fmt.Println("Esto no es Windows")
	} else {
		fmt.Println("Esto es Windows, es ", os)
	}

	switch os := runtime.GOOS; os {
	case "darwin1":
		fmt.Println("Esto es Darwin")
	case "Windows":
		fmt.Println("Esto es Windows")
	default:
		fmt.Printf("%s \n", os)
	}*/

	/*numero, mensaje := ejercicios.ConvierteaInt("19")
	fmt.Println(numero)
	fmt.Println(mensaje)*/

	// teclado.IngresoNumeros()

	// iteraciones.Iterar()
	// texto := ejercicios.CrearTablaNumero()

	// files.GrabaTabla()
	// files.SumaTabla()
	// files.LeoArchivo()
	// files.LeoArchivo2()
	// funciones.Calculos()
	// funciones.LLamarClosure()
	// funciones.Exponencia(2)
	// arreglos_slices.MuestroArreglos()
	// arreglos_slices.MuestroSlice()
	// arreglos_slices.Capacidad()
	// mapas.MostrarMapas()
	// users.AltaUsuario()
	// rufino := new(modelos.Hombre)
	// any := new(modelos.Mujer)
	// e.HumanosRespirando(rufino)
	// e.HumanosRespirando(any)
	// dp.VemosDefer()
	// dp.EjemploPanic()

	canal1 := make(chan bool)
	go goroutines.MiNombreLento("Rufino Bernedo", canal1)
	fmt.Println("Estoy aqui")

	defer func() {
		<-canal1
	}()
	// <-canal1

}
