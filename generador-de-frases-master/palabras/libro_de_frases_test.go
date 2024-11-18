package palabras_test

import (
	"testing"

	"github.com/nickrisaro/generador-de-frases/palabras"
)

func TestAlAgregarUnaFraseDeDosPalabrasSeGeneraUnaListaConFrecuenciasDePalabras(t *testing.T) {

	// Inicialización
	libroDeFrases := palabras.NewLibroDeFrases()

	// Operación
	libroDeFrases.AgregarFrase("¿cómo estás?")

	// Validación
	palabrasProbablesComo := libroDeFrases.ObtenerPalabrasQueSiguenA("¿cómo")

	verificarLongitud(palabrasProbablesComo, 1, "¿cómo", t)

	verificarFrecuencia(palabrasProbablesComo, 1, "estás?", t)

	palabrasProbablesEstas := libroDeFrases.ObtenerPalabrasQueSiguenA("estás?")

	verificarLongitud(palabrasProbablesEstas, 0, "estás?", t)

}

func TestAlAgregarUnaFraseDeTresPalabrasSeGeneraUnaListaConFrecuenciasDePalabras(t *testing.T) {

	// Inicialización
	libroDeFrases := palabras.NewLibroDeFrases()

	// Operación
	libroDeFrases.AgregarFrase("Hola, ¿cómo estás?")

	// Validación
	palabrasProbablesHola := libroDeFrases.ObtenerPalabrasQueSiguenA("Hola,")

	verificarLongitud(palabrasProbablesHola, 1, "Hola,", t)

	verificarFrecuencia(palabrasProbablesHola, 1, "¿cómo", t)

	palabrasProbablesComo := libroDeFrases.ObtenerPalabrasQueSiguenA("¿cómo")

	verificarLongitud(palabrasProbablesComo, 1, "¿cómo", t)

	verificarFrecuencia(palabrasProbablesComo, 1, "estás?", t)

	palabrasProbablesEstas := libroDeFrases.ObtenerPalabrasQueSiguenA("estás?")

	verificarLongitud(palabrasProbablesEstas, 0, "estás?", t)

}

func TestAlAgregarDosFrasesSeGeneraUnaListaConFrecuenciasDePalabras(t *testing.T) {

	// Inicialización
	libroDeFrases := palabras.NewLibroDeFrases()

	// Operación
	libroDeFrases.AgregarFrase("Hola, ¿cómo estás?")
	libroDeFrases.AgregarFrase("Hola, ¿cómo va?")

	// Validación
	palabrasProbablesHola := libroDeFrases.ObtenerPalabrasQueSiguenA("Hola,")

	verificarLongitud(palabrasProbablesHola, 1, "Hola,", t)

	verificarFrecuencia(palabrasProbablesHola, 2, "¿cómo", t)

	palabrasProbablesComo := libroDeFrases.ObtenerPalabrasQueSiguenA("¿cómo")

	verificarLongitud(palabrasProbablesComo, 2, "¿cómo", t)

	verificarFrecuencia(palabrasProbablesComo, 1, "estás?", t)

	verificarFrecuencia(palabrasProbablesComo, 1, "va?", t)

	palabrasProbablesEstas := libroDeFrases.ObtenerPalabrasQueSiguenA("estás?")

	verificarLongitud(palabrasProbablesEstas, 0, "estás?", t)

	palabrasProbablesVa := libroDeFrases.ObtenerPalabrasQueSiguenA("va?")

	verificarLongitud(palabrasProbablesVa, 0, "va?", t)
}

func TestAlAgregarUnaFraseDeUnaPalabraNoSeGeneraUnaListaConFrecuenciasDePalabras(t *testing.T) {

	// Inicialización
	libroDeFrases := palabras.NewLibroDeFrases()

	// Operación
	libroDeFrases.AgregarFrase("Hola")

	// Validación
	palabrasSiguientes := libroDeFrases.ObtenerPalabrasQueSiguenA("Hola")

	verificarLongitud(palabrasSiguientes, 0, "Hola", t)
}

func TestSePuedenObtenerLasDistintasPalabrasQueDiceLaPersona(t *testing.T) {

	// Inicialización
	libroDeFrases := palabras.NewLibroDeFrases()

	// Operación
	libroDeFrases.AgregarFrase("Hola")

	// Validación
	palabrasQueDice := libroDeFrases.ObtenerPalabrasUtilizadas()

	if len(palabrasQueDice) != 1 {
		t.Logf("Se esperaba 1 palabra pero se encontraron %v", len(palabrasQueDice))
		t.FailNow()
	}

	if palabrasQueDice[0] != "Hola" {
		t.Logf("Se esperaba la palabra 'Hola' pero se obtuvo %v", palabrasQueDice[0])
		t.FailNow()
	}
}

func verificarLongitud(palabrasProbables palabras.FrecuenciaPorPalabra, longitudEsperada int, palabra string, t *testing.T) {
	if len(palabrasProbables) != longitudEsperada {
		t.Logf("Se esperaban %v palabras probables para '%v' pero se obtuvieron %v", longitudEsperada, palabra, len(palabrasProbables))
		t.FailNow()
	}
}

func verificarFrecuencia(palabrasProbables palabras.FrecuenciaPorPalabra, frecuenciaEsperada int, palabraBuscada string, t *testing.T) {

	if _, ok := palabrasProbables[palabraBuscada]; !ok {
		t.Logf("Se esperaba la palabra '%v' pero no está", palabraBuscada)
		t.FailNow()
	}

	if palabrasProbables[palabraBuscada] != frecuenciaEsperada {
		t.Logf("Se esperaba que la palabra '%v' tenga frecuencia %v pero se obtuvo %v", palabraBuscada, frecuenciaEsperada, palabrasProbables[palabraBuscada])
		t.FailNow()
	}
}
