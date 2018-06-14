package main
import(
  // "fmt"
  "github.com/kaarla/HOC-Proyecto2/problema_bombero"
  // "github.com/kaarla/HOC-Proyecto2/problema_bombero/grafica"
  // "strconv"
  // "math/rand"
)

func main() {
    problema_bombero.TotalBomberos = 10
    problema_bombero.BomberosXt = 5
    problema_bombero.HormigasXt = 3
    problema_bombero.Phe = 0.3
    problema_bombero.PheReducion = 0.15
    problema_bombero.Semilla = 547

    fuegoInicial := []int{31, 33, 18}

    // semillas := []int{}

    // for i := 0; i < 1000; i++{
    //   semillas = append(semillas, rand.Int())
      // fmt.Println(semillas)
    // }

    // for _, s := range semillas{
      // fmt.Println("Semilla", s)
      // problema_bombero.Semilla = int64(s)
      problema_bombero.CorreHeuristica("graficas/grafica50.txt", fuegoInicial)
    // }



    // grafica := grafica.GeneraCuadricula(100)
    // grafica.ImprimeGrafica()
    // grafica.ImprimeV()
}
