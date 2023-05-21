package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string)
	paises["Chile"] = "Santiago"
	paises["Mexico"] = "DF"
	fmt.Println(paises)
	fmt.Println(paises["Chile"])

	campeonato := map[string]int{
		"Nuggets": 3,
		"Lakers":  0,
		"Heat":    2,
		"Celtics": 0,
	}

	fmt.Println(campeonato)

	for equipo, puntaje := range campeonato {
		fmt.Printf("Equipo %s, ha ganado %d games \n", equipo, puntaje)
	}

	delete(campeonato, "Celtics")
	fmt.Println(campeonato)

	puntaje, existe := campeonato["Spurs"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe = %t \n", puntaje, existe)

}
