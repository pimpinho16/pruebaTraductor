package factory

import (
	"fmt"
	"strconv"
	"strings"
)

type Binario struct {
}

func NewBinario() Binario {
	return Binario{}
}

func (b Binario) Convertir(texto string) (string, error) {
	var binaryStrings []string

	for _, char := range texto {
		asciiValue := int(char)
		binaryString := strconv.FormatInt(int64(asciiValue), 2)
		binaryStrings = append(binaryStrings, fmt.Sprintf("%08s", binaryString)) // Rellena con ceros a la izquierda para obtener 8 bits
	}
	fmt.Printf("Binary code for %s es %s \n", texto, strings.Join(binaryStrings, " "))
	return strings.Join(binaryStrings, " "), nil
}
