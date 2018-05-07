package main
import(
  "fmt"
  "github.com/kaarla/HOC-Proyecto2/problema_bombero"
  "strconv"
)

func main() {
    fmt.Println("hello world")
    vecindario := problema_bombero.VecindarioCero("grafica12.txt")
    for i := 0; i < len(vecindario.Mapa); i++{
      for j := 0; j < len(vecindario.Mapa); j++{
        fmt.Println(strconv.FormatFloat(vecindario.Mapa[i][j], 'f', 6, 64))
      }
    }
}
