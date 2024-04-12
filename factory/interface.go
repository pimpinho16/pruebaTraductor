package factory

type Traductor interface {
	Convertir(texto string) (string, error)
}
