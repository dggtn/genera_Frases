# Generador de frases

Es un pequeño programa que se entrena con una serie de frases que dijo una persona y luego puede decir frases que podría haber dicho esa persona

## Modo de uso

- Escribir en el archivo frases.txt las frases que dice la persona (una por línea)
- Ejecutar go run generador.go

Cuantas más frases se agreguen mejor será el texto generado

## ¿Cómo funciona?

El libro de frases se alimenta de las frases que dijo la persona, para cada palabra de la frase registra cuales son las palabras que la persona dice luego de esa. 

El predictor toma una de las palabras que dijo la persona, busca todas las palabras que podría haber dicho y elige una al azar, luego repite el proceso empezando por la palabra que eligió al azar.

