package main
import(
  "fmt"
  "github.com/kaarla/HOC-Proyecto2/problema_bombero"
  // "strconv"
)

func main() {
    fmt.Println("hello world")
    vecindario := problema_bombero.VecindarioCero("grafica12.txt")
    fmt.Println(vecindario.Manzanas[1].Estado)
    vecindario.InitFuegoEspecifico(6)
    fmt.Println(vecindario.Manzanas[6].Estado)
    fmt.Println(vecindario.GetIncendiados())
    vecindario.Manzanas[3].SetEstado(1)
    vecindario.PropagaFuego()
    fmt.Println(vecindario.GetIncendiados())
    fmt.Println(vecindario.GetDefendidos())
    fmt.Println(vecindario.GetASalvo())
    fmt.Println(vecindario.GetCandidatos())
}
