package factory

import (
	"fmt"
	"github.com/alwindoss/morse"
	"strings"
)

type Morse struct {
}

func NewMorse() Morse {
	return Morse{}
}

func (m Morse) Convertir(texto string) (string, error) {
	h := morse.NewHacker()
	morseCode, err := h.Encode(strings.NewReader(texto))
	if err != nil {
		return "", err
	}
	fmt.Printf("Morse Code is: %s\n", string(morseCode))
	return string(morseCode), nil
}
