package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	asistenciaSkillers := map[string]bool{

		"Carla":    true,
		"Gonzalo":  false,
		"Diego":    false,
		"Federico": true,
		"Luciana":  true,
		"Julian":   false}

	for nombre, asistencia := range asistenciaSkillers {
		fmt.Printf("Skiller: %s asistio?: %t \n", nombre, asistencia)
	}
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese el nombre del alumno para ver si asistio")

	scanner.Scan()
	count, ok := asistenciaSkillers[scanner.Text()]
	if ok {
		fmt.Printf("El skiller asistio? %t \n", count)
	} else {
		fmt.Println("El skiller no existe")
	}
}
