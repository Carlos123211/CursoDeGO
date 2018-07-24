package main

import "fmt"

func main() {
	/* var nombre, apellido string
	nombre = "Carlos"
	apellido = "Gamarra"
	nombre = "Enrique" */

	nombre, apellido := "Carlos", "Gamarra"
	edad := "21"

	fmt.Println(nombre + " " + apellido + " " + "Edad:" + edad)
	fmt.Println(nombre, apellido, "Edad:"+edad)
}
