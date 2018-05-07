package problema-bomberos

 import(
//   "os"
//   "fmt"
   "strings"
   "io/ioutil"
   "strconv"
 )

type Manzana struct{
  id int
  nombre string
  estado string
  vecinos []Manzana
}

type Vecindario struct{
  manzanas []Manzana
  mapa [][]float64
  // var tamano int
  grado int
}

func NewVecindario(mapa [][]float64) *Vecindario{
  vec := Vecindario{}
  vec.mapa = mapa
  //vec.calculaGrado()
  // vec.tamano =
  return &vec
}

func VecindarioCero(grafica string) Vecindario{
  vecindario := Vecindario{}
  vecindario.mapa = initMapa(grafica)
  //TODO
  vecindario.manzanas = nil
  vecindario.grado = 5
  return vecindario
}

func initMapa(grafica string) [][]float64{
  datos, err := ioutil.ReadFile(grafica)
  check(err)
  var mapa [][]float64
  for i := 0; i < len(datos); i++{
    linea := strings.Split(string(datos[i]), ",")
    for j := 0; j < len(datos); j++{
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
