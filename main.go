package main
import(
  "fmt"
  "os"
  "math/rand"
  "github.com/kaarla/HOC-Proyecto2/problema_bombero/grafica"
  "github.com/kaarla/HOC-Proyecto2/problema_bombero"
)

func main() {


  if(len(os.Args) <= 1){
    fmt.Println("Uso: \n Para ejecutar ACO sobre el Problema del bombero ejecutar:",
     "el comando: \n $ go run main.go problema")
  }else{
    if(os.Args[1] == "grafica"){
      grafica := grafica.GeneraCuadricula(50)
      grafica.ImprimeGrafica("graficas/basica30x30.txt")

      distancias, recorridos := grafica.FloydWarshal()
      distancias.ImprimeGrafica("graficas/distancias30x30.txt")
      recorridos.ImprimeGrafica("graficas/recorridos30x30.txt")


    }else if (os.Args[1] == "problema"){
      grafica := "graficas/distancias30x30.txt"
      trayectorias := "graficas/recorridos30x30.txt"
      problema_bombero.Distancias = problema_bombero.InitMapa(grafica)
      problema_bombero.Trayectorias = problema_bombero.InitMapa(trayectorias)
      cantidadBomberos := []int{5, 6, 7, 8, 9}
      semillas := []int{}
      for i := 1; i < 10; i++{
        semillas = append(semillas, rand.Int())
      }

      //fmt.Println(semillas)

      problema_bombero.TotalBomberos = 30
      problema_bombero.HormigasXt = 1
      problema_bombero.Phe = 0.3
      problema_bombero.PheReducion = 0.15
      problema_bombero.PorSalvar = []int{1,6,9,4,15}
      fuegoInicial := []int{70,89,56}

//       for _, seed := range(semillas){
        for _, bomb := range(cantidadBomberos){
          problema_bombero.Semilla = 13850664//int64(seed)
          problema_bombero.BomberosXt = bomb
          problema_bombero.CorreHeuristica(grafica, fuegoInicial)
          fmt.Println("----------------------------")
        }
//         fmt.Println("\n")
//       }
    }
  }
}
