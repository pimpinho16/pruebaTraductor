package cadena

import (
	"fmt"
	"github.com/alwindoss/morse"
	"strings"
)

type Morse struct {
	Siguiente Cadena
}

func NewMorse(Siguiente Cadena) Morse {
	return Morse{Siguiente: Siguiente}
}

func (m Morse) Convertir(texto string) {
	h := morse.NewHacker()
	morseCode, err := h.Encode(strings.NewReader(texto))
	if err != nil {
		fmt.Errorf("Error obtenido %v", err)
	}
	fmt.Printf("Morse Code is: %s\n", string(morseCode))
	m.Siguiente.Convertir(texto)
}
