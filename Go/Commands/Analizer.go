package Commands

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var unit string
var path string
var fit string
var size string
var lines_ []string
var name string

func Line_Comand(line_command__ string) {
	line_command := strings.Split(line_command__, " ")
	Identifier_Command(line_command)
}

func Identifier_Command(array_command []string) {
	command_ := strings.ToLower(array_command[0])
	if command_ == "mkdisk" {
		for i := 1; i < len(array_command); i++ {
			separated := strings.Split(array_command[i], "=")
			parameter := strings.ToLower(separated[0])
			if parameter == "-unit" {
				unit = separated[1]
			} else if parameter == "-path" {
				replacer := strings.NewReplacer("\"", "")
				path = replacer.Replace(separated[1])
			} else if parameter == "-fit" {
				fit = separated[1]
			} else if parameter == "-size" {
				size = separated[1]
			}
		}
		Create_Disk(unit, path, fit, size)
	} else if command_ == "rmdisk" {
		separated := strings.Split(array_command[1], "=")
		replacer := strings.NewReplacer("\"", "")
		path = replacer.Replace(separated[1])
		Delete_Disk(path)
	} else if command_ == "fdisk" {
		for i := 1; i < len(array_command); i++ {
			separated := strings.Split(array_command[i], "=")
			parameter := strings.ToLower(separated[0])
			if parameter == "-name" {
				name = separated[1]
			} else if parameter == "-path" {
				replacer := strings.NewReplacer("\"", "")
				path = replacer.Replace(separated[1])
			} else if parameter == "-size" {
				size = separated[1]
			}

		}
	} else if command_ == "exec" {
		separated := strings.Split(array_command[1], "=")
		replacer := strings.NewReplacer("\"", "")
		path = replacer.Replace(separated[1])
		file, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if scanner.Text() == "\n" || scanner.Text() == "" || scanner.Text() == " " {
			} else {
				time.Sleep(1 * time.Second)
				Line_Comand(scanner.Text())
			}
		}
	} else if command_ == "pause" {
		var pause string
		fmt.Println(yellow + "Presione entrer para continuar: " + closing)
		fmt.Scanln(&pause)
	} else {
		separated := strings.Split(command_, "")
		if separated[0] == "#" {
			var comment string
			for i := 0; i < len(array_command); i++ {
				comment += array_command[i] + " "
			}
			fmt.Println(white + comment + closing)
			comment = ""
		} else {
			fmt.Println(red + "Error: Comando no reconocido." + closing)
		}
	}
}
