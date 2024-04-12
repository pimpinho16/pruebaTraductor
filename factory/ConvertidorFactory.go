package factory

import "errors"

type Factory struct {
	Morse   Traductor
	Binario Traductor
}

func NewConvertidorFactory(Morse Traductor, Binario Traductor) Factory {
	return Factory{
		Morse:   Morse,
		Binario: Binario,
	}
}

func (factory Factory) ConverterFactory(opcion string) (Traductor, error) {
	if opcion == "morse" {
		return factory.Morse, nil
	} else if opcion == "binario" {
		return factory.Binario, nil
	} else {
		return nil, errors.New("function not implemented")
	}
}
