package main
import(
  // "fmt"
  "github.com/kaarla/HOC-Proyecto2/problema_bombero"
  // "github.com/kaarla/HOC-Proyecto2/problema_bombero/grafica"
  // "strconv"
)

func main() {
    problema_bombero.TotalBomberos = 10
    problema_bombero.BomberosXt = 5
    problema_bombero.HormigasXt = 3
    problema_bombero.Phe = 0.3
    problema_bombero.PheReducion = 0.15
    problema_bombero.Semilla = 4845

    fuegoInicial := []int{31, 33, 18, 50, 991, 500, 628, 300}

    problema_bombero.CorreHeuristica("graficas/grafica1000.txt", fuegoInicial)


    // grafica := grafica.GeneraCuadricula(100)
    // grafica.ImprimeGrafica()
    // grafica.ImprimeV()
}
