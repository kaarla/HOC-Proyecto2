package main
import(
  "fmt"
  "os"
  "github.com/kaarla/HOC-Proyecto2/util"
  "github.com/kaarla/HOC-Proyecto2/problema_bombero/grafica"
  "github.com/kaarla/HOC-Proyecto2/problema_bombero"
)

func main() {


  if(len(os.Args) <= 1){
    fmt.Println("Hola, mundo")
  }else{
    if(os.Args[1] == "grafica"){
      grafica := grafica.GeneraCuadricula(10)
      grafica.ImprimeGrafica("graficas/basica10x10.txt")

      distancias, recorridos := grafica.FloydWarshal()
      distancias.ImprimeGrafica("graficas/distancias10x10.txt")
      recorridos.ImprimeGrafica("graficas/recorridos10x10.txt")


    }else if (os.Args[1] == "problema"){
      fuegoInicial := []int{}
      grafica := ""

      problema_bombero.TotalBomberos = 20
      problema_bombero.BomberosXt = 15
      problema_bombero.HormigasXt = 1
      problema_bombero.Phe = 0.3
      problema_bombero.PheReducion = 0.15
      problema_bombero.Semilla =  93558
      problema_bombero.PorSalvar = []int{40, 50}
      grafica = "graficas/distancias10x10.txt"
      trayectorias := "graficas/recorridos10x10.txt"
      problema_bombero.Distancias = problema_bombero.InitMapa(grafica)
      problema_bombero.Trayectorias = problema_bombero.InitMapa(trayectorias)
      fuegoInicial = []int{56}
      problema_bombero.CorreHeuristica(grafica, fuegoInicial)

    }else if (os.Args[1] == "arbol"){
      grafica := grafica.GeneraCuadricula(6)
      distancias, recorridos := grafica.FloydWarshal()
      recorridos.ImprimeGrafica("graficas/recorridos9x9.txt")

      arbol := util.CreaArbol(distancias.Nodos, 4, 9)
      fmt.Println(arbol.Elementos)
      fmt.Println(arbol.GetTrayectoria(8))
    }
  }
}
