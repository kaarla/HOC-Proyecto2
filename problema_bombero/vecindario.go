package problema_bombero

 import(
//   "os"
  // "fmt"
   "strings"
   "io/ioutil"
   "strconv"
 )

type Manzana struct{
  Id int
  Nombre string
  Estado string
  Vecinos []Manzana
}

type Vecindario struct{
  Manzanas []Manzana
  Mapa [][]float64
  // var tamano int
  Grado int
}

func NewVecindario(mapa [][]float64) *Vecindario{
  vec := Vecindario{}
  vec.Mapa = mapa
  //vec.calculaGrado()
  // vec.tamano =
  return &vec
}

func VecindarioCero(grafica string) Vecindario{
  vecindario := Vecindario{}
  vecindario.Mapa = initMapa(grafica)
  //TODO
  vecindario.Manzanas = nil
  vecindario.Grado = 5
  return vecindario
}

func initMapa(grafica string) [][]float64{
  datos, err := ioutil.ReadFile(grafica)
  check(err)
  lineas := strings.Split(string(datos), "\n")

  var mapa [][]float64 = make([][]float64, len(lineas) - 1)
  for k := range mapa{
    mapa[k] = make([]float64, len(lineas) - 1)
  }

  for i := 0; i < len(lineas) - 1; i++{
    linea := strings.Split(string(lineas[i]), ",")
    for j := 0; j < len(lineas) - 1; j++{
      num, err := strconv.ParseFloat(linea[j], 64)
      check(err)
      mapa[i][j] = num
    }
  }
  return mapa
}


func check(e error){
  if e != nil{
    panic(e)
  }
}
