package main
import(
  "fmt"
  "os"
  "github.com/kaarla/HOC-Proyecto2/problema_bombero/grafica"
  "github.com/kaarla/HOC-Proyecto2/problema_bombero"
)

func main() {


  if(len(os.Args) <= 1){
    fmt.Println("Hola, mundo")
  }else{
    if(os.Args[1] == "grafica"){
      grafica := grafica.GeneraCuadricula(6)
      grafica.ImprimeGrafica("graficas/basica9x9.txt")
      //grafica.ImprimeV()

      distancias, recorridos := grafica.FloydWarshal()
      distancias.ImprimeGrafica("graficas/distancias9x9.txt")
      recorridos.ImprimeGrafica("graficas/recorridos9x9.txt")
      fmt.Println(distancias.GetModa(0))

    }else if (os.Args[1] == "problema"){
      fuegoInicial := []int{}
      grafica := ""

      problema_bombero.TotalBomberos = 20
      problema_bombero.BomberosXt = 3
      problema_bombero.HormigasXt = 20
      problema_bombero.Phe = 0.3
      problema_bombero.PheReducion = 0.15
      problema_bombero.Semilla =  54565
      grafica = "graficas/grafica1000.txt"
      fuegoInicial = []int{31, 33, 18, 20}
      problema_bombero.CorreHeuristica(grafica, fuegoInicial)

    }else if (os.Args[1] == "arbol"){
      grafica := grafica.GeneraCuadricula(6)
      distancias, recorridos := grafica.FloydWarshal()
      recorridos.ImprimeGrafica("graficas/recorridos9x9.txt")

      arbol := problema_bombero.CreaArbol(distancias.Nodos, 4, 9)
      fmt.Println(arbol.Elementos)
      fmt.Println(arbol.GetTrayectoria(8))
    }
  }
}
