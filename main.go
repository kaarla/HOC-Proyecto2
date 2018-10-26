package main
import(
  "os"
  "github.com/kaarla/HOC-Proyecto2/problema_bombero"
  "fmt"
)

func main() {

  grafica := ""
  fuegoInicial := []int{}

    if (len(os.Args) <= 1){
      problema_bombero.TotalBomberos = 20
      problema_bombero.BomberosXt = 6
      problema_bombero.HormigasXt = 3
      problema_bombero.Phe = 0.3
      problema_bombero.PheReducion = 0.15
      problema_bombero.Semilla =  54565
      grafica = "graficas/grafica100.txt"
      fuegoInicial = []int{31, 33, 18, 20}

    }else{

    switch os.Args[1] {
    case "50":
      problema_bombero.TotalBomberos = 12
      problema_bombero.BomberosXt = 5
      problema_bombero.HormigasXt = 3
      problema_bombero.Phe = 0.3
      problema_bombero.PheReducion = 0.15
      problema_bombero.Semilla =  547
      grafica = "graficas/grafica50.txt"
      fuegoInicial = []int{31, 33, 18}

    case "1000-a": \\bordes
      problema_bombero.TotalBomberos = 50
      problema_bombero.BomberosXt = 8
      problema_bombero.HormigasXt = 3
      problema_bombero.Phe = 0.3
      problema_bombero.PheReducion = 0.15
      problema_bombero.Semilla = 8710526160774049443
      fuegoInicial = []int{31, 33, 6, 15, 900}
      grafica = "graficas/grafica1000.txt"

    case "1000-b": \\centro
      problema_bombero.TotalBomberos = 50
      problema_bombero.BomberosXt = 12
      problema_bombero.HormigasXt = 3
      problema_bombero.Phe = 0.3
      problema_bombero.PheReducion = 0.15
      problema_bombero.Semilla = 4845
      fuegoInicial = []int{495, 546, 583, 631, 577}
      grafica = "graficas/grafica1000.txt"

    case "":
      fmt.Println("ssss")
    }
  }

    problema_bombero.CorreHeuristica(grafica, fuegoInicial)

}
