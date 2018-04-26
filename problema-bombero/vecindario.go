package bomberos

import(
  "os"
  "fmt"
  "strings"
)

type Manzana struct{
  var estado string
  var id int
  var nombre string
  var vecinos []Nodo
}

type Vecindario struct{
  var manzanas []Manzana
  var mapa [][]float64
  // var tamano int
  var grado int
}

func NewVecindario(mapa [][]float64) *Vecindario{
  vec = Vecindario{}
  vec.mapa = mapa
  vec.calculaGrado()
  // vec.tamano =
  return &vec
}
