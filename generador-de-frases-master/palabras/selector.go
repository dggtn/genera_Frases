package palabras

import (
	"math/rand"
	"time"
)

// Selector quién se encarga de elegir la forma de seleccionar la siguiente palabra
type Selector interface {
	ElegirEntre(primeraPalabra string, primeraFrecuencia int, segundaPalabra string, segundaFrecuencia int) (string, int)
}

// SelectorMaximaFrecuencia selector que elige la palabra de máxima frecuencia
type SelectorMaximaFrecuencia struct {
}

// SelectorPonderado selector que prioriza la palabra de máxima frecuencia
type SelectorPonderado struct {
}

// NewSelectorMaximaFrecuencia crea un nuevo selector de máxima frecuencia
func NewSelectorMaximaFrecuencia() *SelectorMaximaFrecuencia {
	return new(SelectorMaximaFrecuencia)
}

// ElegirEntre elige la palabra con mayor frecuencia
func (s *SelectorMaximaFrecuencia) ElegirEntre(primeraPalabra string, primeraFrecuencia int, segundaPalabra string,
	segundaFrecuencia int) (string, int) {

	if segundaFrecuencia > primeraFrecuencia {
		return segundaPalabra, segundaFrecuencia
	}

	return primeraPalabra, primeraFrecuencia
}

// NewSelectorPonderado crea un nuevo selector ponderado
func NewSelectorPonderado() *SelectorPonderado {
	return new(SelectorPonderado)
}

// ElegirEntre elige la palabra dando más prioridad a la que tiene más frecuencia
func (s *SelectorPonderado) ElegirEntre(primeraPalabra string, primeraFrecuencia int, segundaPalabra string,
	segundaFrecuencia int) (string, int) {

	rand.Seed(time.Now().UnixNano())

	palabras := []string{}

	for i := 0; i < primeraFrecuencia; i++ {

		palabras = append(palabras, primeraPalabra)
	}

	for i := 0; i < segundaFrecuencia; i++ {

		palabras = append(palabras, segundaPalabra)
	}

	palabraElegida := palabras[rand.Intn(len(palabras))]

	if palabraElegida == primeraPalabra {
		return primeraPalabra, primeraFrecuencia
	}

	return segundaPalabra, segundaFrecuencia

}
