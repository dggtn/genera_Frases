

  <h1>Bot Generador de Frases Célebres de Borges</h1>
  <p><strong>Cursando: Proyecto de Programación en Go</strong></p>

  <h2>Descripción del Proyecto</h2>
  <p>Este proyecto consiste en un bot que genera frases célebres del escritor argentino <strong>Jorge Luis Borges</strong>. La aplicación está escrita en <strong>Go</strong> y permite obtener una cita aleatoria de Borges cada vez que se ejecuta.</p>
  <p>El objetivo es familiarizarse con el lenguaje Go y cómo trabajar con estructuras básicas de datos, como arreglos o slices, y funciones para manipularlas.</p>

  <h2>Tecnologías Utilizadas</h2>
  <ul>
      <li><strong>Go</strong>: Lenguaje de programación utilizado para desarrollar el bot.</li>
  </ul>

  <h2>Características del Proyecto</h2>
  <ul>
      <li><strong>Generación de frases aleatorias:</strong> El bot selecciona una frase célebre aleatoria de Borges de una lista predefinida de citas.</li>
      <li><strong>Interfaz sencilla:</strong> El bot simplemente imprime la frase en la consola.</li>
      <li><strong>Código simple y eficiente:</strong> El proyecto utiliza las estructuras y funciones básicas de Go.</li>
  </ul>

  <h2>Instrucciones para Ejecutar el Proyecto</h2>

  <h3>Requisitos Previos</h3>
  <ul>
      <li><strong>Go</strong> instalado en tu sistema. Si no tienes Go instalado, puedes descargarlo desde <a href="https://golang.org/dl/" target="_blank">golang.org</a>.</li>
  </ul>

  <h3>Pasos para Ejecutar el Proyecto:</h3>
  <ol>
      <li><strong>Clona el repositorio</strong>:
          <pre><code>git clone https://github.com/tu-usuario/bot-frases-borges.git</code></pre>
      </li>
      <li><strong>Navega al directorio del proyecto</strong>:
          <pre><code>cd bot-frases-borges</code></pre>
      </li>
      <li><strong>Ejecuta el programa</strong>:
          <pre><code>go run main.go</code></pre>
      </li>
      <li>El bot imprimirá una frase célebre de Borges en la consola.</li>
  </ol>

  <h2>Código de Ejemplo</h2>
  <p>A continuación, te mostramos un ejemplo de cómo está estructurado el código en Go:</p>

  <pre><code>
// main.go
package main

import (
  "fmt"
  "math/rand"
  "time"
)

// Lista de frases célebres de Borges
var frases = []string{
  "“No hay victoria sin derrota, ni derrota sin victoria.” - Jorge Luis Borges",
  "“La duda es uno de los nombres de la inteligencia.” - Jorge Luis Borges",
  "“Uno está enamorado cuando se da cuenta de que otra persona es única.” - Jorge Luis Borges",
  "“El futuro es el presente de otro.” - Jorge Luis Borges",
  "“La biblioteca es un lugar perfecto: no tiene tiempo ni espacio.” - Jorge Luis Borges",
}

// Función para generar una frase aleatoria
func obtenerFraseAleatoria() string {
  rand.Seed(time.Now().UnixNano())
  indice := rand.Intn(len(frases))
  return frases[indice]
}

func main() {
  fmt.Println("Frase célebre de Borges:")
  fmt.Println(obtenerFraseAleatoria())
}
  </code></pre>

  <h2>Contribución</h2>
  <p>Si deseas contribuir a este proyecto, por favor sigue estos pasos:</p>
  <ol>
      <li>Haz un fork del repositorio.</li>
      <li>Crea una nueva rama (<code>git checkout -b feature/nueva-funcionalidad</code>).</li>
      <li>Realiza tus cambios y haz commits.</li>
      <li>Abre un pull request describiendo tus cambios.</li>
  </ol>

  <h2>Licencia</h2>
  <p>Este proyecto está bajo la licencia <strong>MIT</strong>. Consulta el archivo LICENSE para más detalles.</p>

</body>
</html>
