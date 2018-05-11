package problema_bombero

 import(
  "fmt"
   "strings"
   "io/ioutil"
   "strconv"
 )

type Manzana struct{
  Id int
  // Nombre string
  Estado int
  Vecinos []int
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
  vecindario.Manzanas = nil
  vecindario.Grado = 5
  vecindario.initManzanas()
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

func initManzana(id int) Manzana{
  manzana := Manzana{}
  manzana.Id = id
  manzana.Estado = 0
  manzana.Vecinos = nil
  return manzana
}

func (vecindario *Vecindario) initManzanas(){
  mapa := vecindario.Mapa
  var manzanas []Manzana = make ([]Manzana, len(mapa))
  for i := 0; i < len(mapa); i++{
    vecinos := []int{}
    manzanas[i] = initManzana(i)
    for j := 0; j < len(mapa); j++{
      if(mapa[i][j] > 0.0 && i != j){
        vecinos = append(vecinos, j)
      }
    }
    manzanas[i].Vecinos = vecinos
    vecindario.Manzanas = manzanas
  }
}

// TODO: fuego random
// func (vecindario *Vecindario) initFuegoRandom(semilla int){
// }

func (manzana *Manzana) SetEstado(estado int){
  manzana.Estado = estado
}

func (vecindario *Vecindario) InitFuegoEspecifico(manzana int){
    vecindario.Manzanas[manzana].SetEstado(2)
}

func (vecindario *Vecindario) PropagaFuego(){
  incendiados := vecindario.GetIncendiados()
  for i := 0; i < len(incendiados); i++{
    v := vecindario.Manzanas[incendiados[i]].Vecinos
    for j := 0; j < len(v); j++{
      m := vecindario.Manzanas[v[j]]
      if(m.Estado == 0){
        vecindario.Manzanas[v[j]].SetEstado(2)
      }
    }
  }
}

func (vecindario *Vecindario) GetIncendiados() []int{
  res := []int{}
  for i := 0; i < len(vecindario.Manzanas); i++{
    if(vecindario.Manzanas[i].Estado == 2){
      res = append(res, i)
    }
  }
  return res
}

func (vecindario *Vecindario) GetASalvo() []int{
  res := []int{}
  i := 0
  for _, b := range vecindario.Manzanas{
    if(b.Estado == 0){
      res = append(res, i)
    }
    i++
  }
  return res
}

func (vecindario *Vecindario) GetDefendidos() []int{
  res := []int{}
  for i := 0; i < len(vecindario.Manzanas); i++{
    if(vecindario.Manzanas[i].Estado == 1){
      res = append(res, i)
    }
  }
  return res
}


func (vecindario *Vecindario) GetCandidatos() []int{
  incendiados := vecindario.GetIncendiados()
  candidatos := []int{}
  for i := 0; i < len(incendiados); i++{
    v := vecindario.Manzanas[incendiados[i]].Vecinos
    for j := 0; j < len(v); j++{
      m := vecindario.Manzanas[v[j]]
      if(m.Estado == 0 && !Contiene(candidatos, v[j])){
        candidatos = append(candidatos, v[j])
      }
    }
  }
  fmt.Println("")
  return candidatos
}

func Contiene(a []int, e int) bool{
  for _, b := range a{
    if b == e{
      return true
    }
  }
  return false
}

func(vecindario *Vecindario) Copia() Vecindario{
  copia := Vecindario{}
  copia.Mapa = vecindario.Mapa
  copia.Manzanas = vecindario.Manzanas
  copia.Grado = vecindario.Grado
  return copia
}

func check(e error){
  if e != nil{
    panic(e)
  }
}

func (vecindario *Vecindario) Evalua(numBomberos int) float64{
  quemados := float64(len(vecindario.GetIncendiados()))
  defendidos := float64(len(vecindario.GetDefendidos()))
  return (quemados / float64(len(vecindario.Manzanas))) * (defendidos / float64(numBomberos))
}
