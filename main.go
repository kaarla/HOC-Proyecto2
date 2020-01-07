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
    fmt.Println("Uso: \n Para ejecutar ACO sobre el Problema del bombero correr",
     "el comando: \n $ go run main.go problema")
  }else{
    if(os.Args[1] == "grafica"){
      grafica := grafica.GeneraCuadricula(105)
      grafica.ImprimeGrafica("graficas/basica30x30.txt")

      distancias, recorridos := grafica.FloydWarshal()
      distancias.ImprimeGrafica("graficas/distancias30x30.txt")
      recorridos.ImprimeGrafica("graficas/recorridos30x30.txt")


    }else if (os.Args[1] == "problema"){
      fuegoInicial := []int{}
      grafica := ""

      problema_bombero.TotalBomberos = 20
      problema_bombero.BomberosXt = 5
      problema_bombero.HormigasXt = 3
      problema_bombero.Phe = 0.3
      problema_bombero.PheReducion = 0.15
      problema_bombero.Semilla =  9355888
      problema_bombero.PorSalvar = []int{1, 10}

      grafica = "graficas/distancias10x10.txt"
      trayectorias := "graficas/recorridos10x10.txt"
      problema_bombero.Distancias = problema_bombero.InitMapa(grafica)
      problema_bombero.Trayectorias = problema_bombero.InitMapa(trayectorias)
      fuegoInicial = []int{36, 66}
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
