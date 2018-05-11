package main
import(
  // "fmt"
  "github.com/kaarla/HOC-Proyecto2/problema_bombero"
  // "strconv"
)

func main() {
    problema_bombero.TotalBomberos = 10
    problema_bombero.BomberosXt = 2
    problema_bombero.HormigasXt = 3
    problema_bombero.Phe = 0.3
    problema_bombero.PheReducion = 0.15
    problema_bombero.Semilla = 0

    problema_bombero.CorreHeuristica("grafica12.txt")

    // fmt.Println("hello world")
    // vecindario := problema_bombero.VecindarioCero("grafica12.txt")
    // fmt.Println(vecindario.Manzanas[1].Estado)
    // vecindario.InitFuegoEspecifico(6)
    // escenario1 := problema_bombero.InitEscenario(vecindario)
    // fmt.Println(vecindario.Manzanas[6].Estado)
    // fmt.Println(vecindario.GetIncendiados())
    // vecindario.Manzanas[3].SetEstado(1)
    // vecindario.PropagaFuego()
    // fmt.Println(vecindario.GetIncendiados())
    // fmt.Println(vecindario.GetDefendidos())
    // fmt.Println(vecindario.GetASalvo())
    // fmt.Println(vecindario.GetCandidatos())
    // fmt.Println("eval ", vecindario.Evalua(10))
    // escenario := problema_bombero.InitEscenario(vecindario)
    // fmt.Println("son iguales?", (escenario.Ve.Manzanas[0].Id == escenario1.Ve.Manzanas[0].Id))
    // fmt.Println(escenario.Eval)
    // // hormiga := problema_bombero.InitHormiga(1, escenario)
    // hormiga2 := problema_bombero.InitHormiga(2, escenario1)
    // hormiga2.Trayecto = append(hormiga2.Trayecto, escenario)
    // hormiga2.AvanzaHormiga()
    // // fmt.Println(hormiga)
    // // fmt.Println("actual", hormiga2.Actual)
    // // fmt.Println("siguiente", hormiga2.Siguiente)
    // // fmt.Println("Trayecto", hormiga2.Trayecto)
    // fmt.Println("Costo", hormiga2.CalculaCosto())
}
