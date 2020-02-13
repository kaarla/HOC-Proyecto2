package main
import(
  "database/sql"
  "fmt"
  "os"
  "github.com/kaarla/HOC-Proyecto2/problema_bombero/grafica"
  "github.com/kaarla/HOC-Proyecto2/problema_bombero"
  _ "github.com/mattn/go-sqlite3"
)



func check(e error) {
    if e != nil {
        fmt.Println(e)
        panic(e)
    }
}

func main() {
  if(len(os.Args) <= 1){
    fmt.Println("Uso: \n Para ejecutar ACO sobre el Problema del bombero correr",
     "el comando: \n $ go run main.go problema")
  }else{
    if(os.Args[1] == "grafica"){
      var err error
      grafica.GraphDB, err = sql.Open("sqlite3", "databases/Grafica3x3.db")
      check(err)
      errS := grafica.GeneraBaseCuadricula(4)
      fmt.Printf("errS: %d ", errS)
      // grafica := grafica.GeneraCuadricula(25)
      // grafica.ImprimeGrafica("graficas/basica5x5.txt")

      distancias, recorridos := grafica.FloydWarshal()
      distancias.ImprimeGrafica("graficas/distancias75x75.txt")
      recorridos.ImprimeGrafica("graficas/recorridos75x75.txt")


    }else if (os.Args[1] == "problema"){
      fuegoInicial := []int{}
      grafica := ""

      problema_bombero.TotalBomberos = 20
      problema_bombero.BomberosXt = 9
      problema_bombero.HormigasXt = 6
      problema_bombero.Phe = 0.3
      problema_bombero.PheReducion = 0.15
      problema_bombero.Semilla =  9355
      problema_bombero.PorSalvar = []int{300, 400, 850, 410}

      grafica = "graficas/distancias30x30.txt"
      trayectorias := "graficas/recorridos30x30.txt"
      problema_bombero.Distancias = problema_bombero.InitMapa(grafica)
      problema_bombero.Trayectorias = problema_bombero.InitMapa(trayectorias)
      fuegoInicial = []int{1, 800, 350, 389, 246, 288, 236, 741, 896, 200}
      problema_bombero.CorreHeuristica(grafica, fuegoInicial)

    }
  }
}
