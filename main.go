package main
import(
  // "fmt"
  "github.com/kaarla/HOC-Proyecto2/problema_bombero"
  // "github.com/kaarla/HOC-Proyecto2/problema_bombero/grafica"
  // "strconv"
)

func main() {
    problema_bombero.TotalBomberos = 10
    problema_bombero.BomberosXt = 4
    problema_bombero.HormigasXt = 3
    problema_bombero.Phe = 0.3
    problema_bombero.PheReducion = 0.15
    problema_bombero.Semilla = 50

    problema_bombero.CorreHeuristica("grafica50.txt")


    // grafica := grafica.GeneraCuadricula(15)
    // grafica = grafica.DiagonalesRandom(5)
    // grafica.ImprimeGrafica()
    // grafica.ImprimeV()
}
