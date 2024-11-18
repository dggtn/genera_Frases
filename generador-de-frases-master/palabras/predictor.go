package palabras

import (
	"math/rand"
	"time"
)

// Predictor es el encargado de generar frases que podría decir un ser humano
type Predictor struct {
	libroDeFrases *LibroDeFrases
	selector      Selector
}

// NewPredictor construye un nuevo Predictor
func NewPredictor(libroDeFrases *LibroDeFrases, selector Selector) *Predictor {

	predictor := new(Predictor)
	predictor.libroDeFrases = libroDeFrases
	predictor.selector = selector

	return predictor
}

// GenerarFrase genera una frase que comienza con alguna de las palabras que la persona dice
func (p *Predictor) GenerarFrase(cantidadDePalabras int) string {

	primeraPalabra := p.obtenerPalabraAlAzar()
	return p.GenerarFraseAPartirDe(primeraPalabra, cantidadDePalabras)
}

func (p *Predictor) obtenerPalabraAlAzar() string {

	rand.Seed(time.Now().UnixNano())

	palabras := p.libroDeFrases.ObtenerPalabrasUtilizadas()

	return palabras[rand.Intn(len(palabras))]

}

// GenerarFraseAPartirDe genera una frase que comienza con la palabra primeraPalabra y tiene como máximo cantidadDePalabras palabras
func (p *Predictor) GenerarFraseAPartirDe(primeraPalabra string, cantidadDePalabras int) string {

	palabrasDeLaFrase := 1
	frase := primeraPalabra
	siguientesPalabras := p.libroDeFrases.ObtenerPalabrasQueSiguenA(primeraPalabra)

	for palabrasDeLaFrase < cantidadDePalabras {

		siguientePalabra := p.buscarPalabraConMayorFrecuencia(siguientesPalabras)

		frase = frase + " " + siguientePalabra
		palabrasDeLaFrase++

		siguientesPalabras = p.libroDeFrases.ObtenerPalabrasQueSiguenA(siguientePalabra)
	}

	return frase
}

func (p *Predictor) buscarPalabraConMayorFrecuencia(siguientesPalabras FrecuenciaPorPalabra) string {

	frecuenciaMaxima := 0
	palabra := ""

	for k, v := range siguientesPalabras {
		palabra, frecuenciaMaxima = p.selector.ElegirEntre(palabra, frecuenciaMaxima, k, v)
	}

	return palabra

}
