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
    fmt.Println("Uso: \n Para ejecutar ACO sobre el Problema del bombero ejecuta",
     "el comando: \n $ go run main.go problema")
  }else{
    //sql.Open("sqlite3", "databases/Grafica3x3.db")

    grafica.GraphDB, _ = sql.Open("sqlite3", "databases/Grafica3x3.db")

    if(os.Args[1] == "grafica"){
      errS := grafica.GeneraBaseCuadricula(1)
      fmt.Printf("errS: %d ", errS)
      grafica.FloydWarshal()

    }else if (os.Args[1] == "problema"){
      fuegoInicial := []int{}
      // grafica.GraphDB, _ = sql.Open("sqlite3", "databases/Grafica3x3.db")
      // check(error)

      problema_bombero.NumVertices = 9
      problema_bombero.TotalBomberos = 20
      problema_bombero.BomberosXt = 9
      problema_bombero.HormigasXt = 6
      problema_bombero.Phe = 0.3
      problema_bombero.PheReducion = 0.15
      problema_bombero.Semilla =  9355
      problema_bombero.PorSalvar = []int{1}

      fuegoInicial = []int{6}
      problema_bombero.CorreHeuristica(fuegoInicial)

    }
  }
}
