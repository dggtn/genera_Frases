package palabras_test

import (
	"strings"
	"testing"

	"github.com/nickrisaro/generador-de-frases/palabras"
)

func TestGenerarFraseDeDosPalabras(t *testing.T) {

	// Inicialización
	libroDeFrases := palabras.NewLibroDeFrases()
	libroDeFrases.AgregarFrase("¿cómo estás?")
	selector := palabras.NewSelectorMaximaFrecuencia()
	predictor := palabras.NewPredictor(libroDeFrases, selector)

	// Operación
	fraseGenerada := predictor.GenerarFraseAPartirDe("¿cómo", 2)

	// Validación
	fraseEsperada := "¿cómo estás?"

	if fraseGenerada != fraseEsperada {
		t.Logf("Se esperaba '%v' pero se generó '%v'", fraseEsperada, fraseGenerada)
		t.FailNow()
	}

}

func TestGenerarFraseDeTresPalabras(t *testing.T) {

	// Inicialización
	libroDeFrases := palabras.NewLibroDeFrases()
	libroDeFrases.AgregarFrase("Hola, ¿cómo estás?")
	selector := palabras.NewSelectorMaximaFrecuencia()
	predictor := palabras.NewPredictor(libroDeFrases, selector)

	// Operación
	fraseGenerada := predictor.GenerarFraseAPartirDe("Hola,", 3)

	// Validación
	fraseEsperada := "Hola, ¿cómo estás?"

	if fraseGenerada != fraseEsperada {
		t.Logf("Se esperaba '%v' pero se generó '%v'", fraseEsperada, fraseGenerada)
		t.FailNow()
	}

}

func TestGenerarFraseDeTresPalabrasConLasPalabrasMasFrecuentes(t *testing.T) {

	// Inicialización
	libroDeFrases := palabras.NewLibroDeFrases()
	libroDeFrases.AgregarFrase("Hola, ¿qué tal?")
	libroDeFrases.AgregarFrase("Hola, ¿cómo estás?")
	libroDeFrases.AgregarFrase("Hola, ¿cómo va?")
	selector := palabras.NewSelectorMaximaFrecuencia()
	predictor := palabras.NewPredictor(libroDeFrases, selector)

	// Operación
	fraseGenerada := predictor.GenerarFraseAPartirDe("Hola,", 3)

	// Validación
	frasesEsperadas := []string{"Hola, ¿cómo estás?", "Hola, ¿cómo va?"}

	if fraseGenerada != frasesEsperadas[0] && fraseGenerada != frasesEsperadas[1] {
		t.Logf("Se esperaba '%v' pero se generó '%v'", frasesEsperadas, fraseGenerada)
		t.FailNow()
	}

}

func TestGenerarFraseDeTresPalabrasEmpezandoPorUnaPalabraAlAzar(t *testing.T) {

	// Inicialización
	libroDeFrases := palabras.NewLibroDeFrases()
	libroDeFrases.AgregarFrase("¿qué tal?")
	selector := palabras.NewSelectorMaximaFrecuencia()
	predictor := palabras.NewPredictor(libroDeFrases, selector)

	// Operación
	fraseGenerada := predictor.GenerarFrase(2)

	// Validación

	if !strings.HasPrefix(fraseGenerada, "¿qué") && !strings.HasPrefix(fraseGenerada, "tal?") {
		t.Logf("La frase generada no empieza con las palabras esperadas, se generó '%v'", fraseGenerada)
		t.FailNow()
	}

}
