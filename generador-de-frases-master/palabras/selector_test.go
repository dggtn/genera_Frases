package palabras_test

import (
	"testing"

)

func TestElSelectorMaximaFrecuenciaEligeLaPalabraDeMayorFrecuencia(t *testing.T) {

	// Inicialización
	selector := palabras.NewSelectorMaximaFrecuencia()

	primeraPalabra, primeraFrecuencia := "Hola", 3
	segundaPalabra, segundaFrecuencia := "Chau", 2

	// Operación
	palabra, frecuencia := selector.ElegirEntre(primeraPalabra, primeraFrecuencia, segundaPalabra, segundaFrecuencia)

	// Validación
	if palabra != primeraPalabra {
		t.Logf("Se esperaba '%v' pero se obtuvo '%v'", primeraPalabra, palabra)
		t.FailNow()
	}

	if frecuencia != primeraFrecuencia {
		t.Logf("Se esperaba '%v' pero se obtuvo '%v'", primeraFrecuencia, frecuencia)
		t.FailNow()
	}
}

func TestElSelectorMaximaFrecuenciaEligeLaPrimeraPalabraSiTienenIgualFrecuencia(t *testing.T) {

	// Inicialización
	selector := palabras.NewSelectorMaximaFrecuencia()

	primeraPalabra, primeraFrecuencia := "Hola", 3
	segundaPalabra, segundaFrecuencia := "Chau", 3

	//Operación
	palabra, frecuencia := selector.ElegirEntre(primeraPalabra, primeraFrecuencia, segundaPalabra, segundaFrecuencia)

	// Validación
	if palabra != primeraPalabra {
		t.Logf("Se esperaba '%v' pero se obtuvo '%v'", primeraPalabra, palabra)
		t.FailNow()
	}

	if frecuencia != primeraFrecuencia {
		t.Logf("Se esperaba '%v' pero se obtuvo '%v'", primeraFrecuencia, frecuencia)
		t.FailNow()
	}
}

func TestElSelectorPonderadoEligeMasFrecuentementeLaPalabraConMayorFrecuencia(t *testing.T) {

	// Inicialización
	selector := palabras.NewSelectorPonderado()

	primeraPalabra, primeraFrecuencia := "Hola", 9
	segundaPalabra, segundaFrecuencia := "Chau", 1

	cantidadDeVecesQueEligeHola, cantidadDeVecesQueEligeChau := 0, 0

	// Operación
	for i := 0; i < 100; i++ {
		palabraElegida, _ := selector.ElegirEntre(primeraPalabra, primeraFrecuencia, segundaPalabra, segundaFrecuencia)

		if palabraElegida == "Hola" {
			cantidadDeVecesQueEligeHola++
		} else {
			cantidadDeVecesQueEligeChau++
		}
	}

	// Validación
	if cantidadDeVecesQueEligeHola < 80 {
		t.Logf("Se obtuvo la palabra 'Hola' con una frecuencia menor a la esperada: %v", cantidadDeVecesQueEligeHola)
		t.FailNow()
	}

	if cantidadDeVecesQueEligeHola > 95 {
		t.Logf("Se obtuvo la palabra 'Hola' con una frecuencia mayor a la esperada: %v", cantidadDeVecesQueEligeHola)
		t.FailNow()
	}

	if cantidadDeVecesQueEligeChau < 5 {
		t.Logf("Se obtuvo la palabra 'Chau' con una frecuencia menor a la esperada: %v", cantidadDeVecesQueEligeChau)
		t.FailNow()
	}

	if cantidadDeVecesQueEligeChau > 20 {
		t.Logf("Se obtuvo la palabra 'Chau' con una frecuencia mayor a la esperada: %v", cantidadDeVecesQueEligeChau)
		t.FailNow()
	}
}
