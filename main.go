package main
import(
  "fmt"
  "github.com/kaarla/HOC-Proyecto2/problema_bombero"
)

func main() {
    fmt.Println("hello world")
    vecindario := problema_bombero.VecindarioCero("grafica12.txt")
    fmt.Println(vecindario.getMapa())
}
