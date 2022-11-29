package main

import (
	f "Conferencia_Archivos/Commands"
	"fmt"
)

var yellow string = "\033[1;33m"
var closing string = "\033[0m"

func main() {
	iteraccion := 1
	for iteraccion < 1000 {
		var first_part string
		var second_part string
		var pause string
		fmt.Print(yellow + ">>" + closing)
		fmt.Scanln(&first_part, &second_part)
		if first_part == "exit" {
			iteraccion = 1000
			fmt.Println(yellow + "Se termino la ejecucion del programa" + closing)
		} else if first_part == "pause" {
			fmt.Println(yellow + "Presione entrer para continuar: " + closing)
			fmt.Scanln(&pause)
		} else {
			f.Line_Comand(first_part + " " + second_part)
			iteraccion += iteraccion
		}
	}
}
