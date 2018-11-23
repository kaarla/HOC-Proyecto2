package main
import(
  // "os"
  "github.com/kaarla/HOC-Proyecto2/problema_bombero/grafica"
)

func main() {

  grafica := grafica.GeneraCuadricula(6)
  grafica.ImprimeGrafica()
  grafica.ImprimeV()

  recorridos := grafica.FloydWarshal()
  recorridos.ImprimeGrafica()

}
