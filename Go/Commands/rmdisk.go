package Commands

import (
	"fmt"
	"os"
)

func Delete_Disk(path string) {
	disk_name := find_last_of(path)
	if path == "" {
		fmt.Println(red + "Error: Falta el parametro obligatorio path." + closing)
	} else {
		fmt.Println(blue + "¿Desea eliminar el disco " + closing + yellow + disk_name + closing + blue + " ?  [Yes/No]" + closing)
		var Opcion string
		fmt.Scanln(&Opcion)
		if Opcion == "yes" {
			err := os.Remove(path)
			if err != nil {
				fmt.Println(red + "Error: No existe el disco " + closing + yellow + disk_name + closing + red + " para ser eliminado." + closing)
			}
			fmt.Println(green + "El disco " + closing + yellow +
				disk_name + closing + green +
				" fue eliminado exitosamente de la ruta " + closing + yellow + path + closing)
		} else if Opcion == "no" {
			fmt.Println(green + "No se eliminara el disco " + closing + yellow + disk_name + closing)
		}
	}
}
