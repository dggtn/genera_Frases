package palabras

import "strings"

// FrecuenciaPorPalabra indica con qué frecuencia un ser humano dice dicha palabra
type FrecuenciaPorPalabra map[string]int

// LibroDeFrases es la estructura encargada de guardar la información de las frase que dice un ser humano
type LibroDeFrases struct {
	frecuencia map[string]FrecuenciaPorPalabra
}

// NewLibroDeFrases crea un libroDeFrases
func NewLibroDeFrases() *LibroDeFrases {
	libroDeFrases := new(LibroDeFrases)
	libroDeFrases.frecuencia = make(map[string]FrecuenciaPorPalabra)
	return libroDeFrases
}

// AgregarFrase le dice al libroDeFrases que un ser humano dice esa frase
func (l *LibroDeFrases) AgregarFrase(frase string) {

	palabrasDeLaFrase := strings.Fields(frase)

	if len(palabrasDeLaFrase) > 1 {

		for i := range palabrasDeLaFrase {

			if i < len(palabrasDeLaFrase)-1 {

				primeraPalabra := palabrasDeLaFrase[i]

				palabrasSiguientes := l.ObtenerPalabrasQueSiguenA(primeraPalabra)

				if palabrasSiguientes == nil {
					palabrasSiguientes = make(FrecuenciaPorPalabra)
					l.frecuencia[primeraPalabra] = palabrasSiguientes
				}

				segundaPalabra := palabrasDeLaFrase[i+1]

				palabrasSiguientes[segundaPalabra]++

			}

		}
	} else {
		l.frecuencia[palabrasDeLaFrase[0]] = make(FrecuenciaPorPalabra)
	}

}

// ObtenerPalabrasQueSiguenA devuelve una lista de palabras que puede decir un ser humano luego de la palabra recibida por parámetro
func (l *LibroDeFrases) ObtenerPalabrasQueSiguenA(palabra string) FrecuenciaPorPalabra {
	return l.frecuencia[palabra]
}

// ObtenerPalabrasUtilizadas devuelve todas las palabras que usa la persona
func (l *LibroDeFrases) ObtenerPalabrasUtilizadas() []string {

	palabras := []string{}

	for k := range l.frecuencia {
		palabras = append(palabras, k)
	}

	return palabras
}
