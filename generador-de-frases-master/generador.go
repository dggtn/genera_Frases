package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

)

func main() {

	libro := palabras.NewLibroDeFrases()

	file, err := os.Open("frases.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		frase := scanner.Text()

		if frase != "" {
			libro.AgregarFrase(frase)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("------------Entrenamiento finalizado------------------")
	selector := palabras.NewSelectorPonderado()
	predictor := palabras.NewPredictor(libro, selector)

	fmt.Println("Voy a predecir una frase")
	fmt.Println("Para terminar poné una letra en lugar de cantidad de palabras")

	for err == nil {

		var palabraInicial, palabras string

		fmt.Print("Cómo empieza?: ")
		fmt.Scanln(&palabraInicial)
		fmt.Print("De cuántas palabras?: ")
		fmt.Scanln(&palabras)

		cantidadDePalabras, err := strconv.Atoi(palabras)

		if err != nil {
			return
		}

		if palabraInicial != "" {
			fmt.Println(predictor.GenerarFraseAPartirDe(palabraInicial, cantidadDePalabras))
		} else {
			fmt.Println(predictor.GenerarFrase(cantidadDePalabras))
		}
	}

}
