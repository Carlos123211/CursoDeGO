package main

import "fmt"

func main() {
	/* var nombre, apellido string
	nombre = "Carlos"
	apellido = "Gamarra"
	nombre = "Enrique" */

	nombre, apellido := "Carlos", "Gamarra222"
	edad := "21"

	fmt.Println(nombre + " " + apellido + " " + "Edad:" + edad)
	fmt.Println(nombre, apellido, "Edad:"+edad)
}
