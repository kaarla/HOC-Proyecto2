package main
import(
  "fmt"
  "github.com/kaarla/HOC-Proyecto2/problema_bombero/grafica"
  "github.com/kaarla/HOC-Proyecto2/problema_bombero"
  "math/rand"
  "os"
  "strconv"
)

func main() {
  graficaFile := "graficas/basica23x23.txt"
  distanciasFile := "graficas/distancias10x10.txt"
  recorridosFile := "graficas/recorridos10x10.txt"

  if(len(os.Args) <= 1){
    fmt.Println("Uso: \n Para ejecutar ACO sobre el Problema del bombero ejecutar:",
     "el comando: \n $ go run main.go problema")
  }else{
    if(os.Args[1] == "grafica"){
      grafica := grafica.GeneraCuadricula(50)
      grafica.ImprimeGrafica(graficaFile)

      distancias, recorridos := grafica.FloydWarshal()
      distancias.ImprimeGrafica(distanciasFile)
      recorridos.ImprimeGrafica(recorridosFile)


    }else if (os.Args[1] == "problema"){
      grafica := distanciasFile
      trayectorias := recorridosFile
      problema_bombero.Distancias = problema_bombero.InitMapa(grafica)
      problema_bombero.Trayectorias = problema_bombero.InitMapa(trayectorias)
      semillas := []int{}
      for i := 1; i < 10; i++{
        semillas = append(semillas, rand.Int())
      }

      problema_bombero.TotalBomberos = 30
      problema_bombero.HormigasXt = 6
      problema_bombero.Phe = 0.3
      problema_bombero.PheReducion = 0.15
      problema_bombero.PorSalvar = []int{1,6,9,4,15}
      fuegoInicial := []int{70,89,56}


      seed, _ := strconv.Atoi(os.Args[2])
      problema_bombero.Semilla = int64(seed)
      problema_bombero.BomberosXt, _ = strconv.Atoi(os.Args[3])
      problema_bombero.CorreHeuristica(grafica, fuegoInicial)
    }
  }
}
