package main

import (
	"fmt"
	"pruebatraductor/cadena"
)

func main() {
	fmt.Println("Ejemplo")

	binario := cadena.NewMBinario()
	morse := cadena.NewMorse(binario)

	morse.Convertir("Alexander")

}
