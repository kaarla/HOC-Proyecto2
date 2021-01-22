package problema_bombero

 import(
  "github.com/kaarla/HOC-Proyecto2/problema_bombero/grafica"
  "fmt"
  "github.com/kaarla/HOC-Proyecto2/util"
  "strings"
  "strconv"
 )

type Vecindario struct{
   Grado int
   Cantidad int
   Incendiados int
   ASalvo int
   Defendidos int
   Manzanas []Manzana
}

type Manzana struct{
  Id int
  Estado int
}

func NewVecindario(mapa [][]int) *Vecindario{
  vec := Vecindario{}
  return &vec
}

func VecindarioCero() Vecindario{
  vecindario := Vecindario{}
  vecindario.Grado = 5
  vecindario.initManzanas()
  return vecindario
}

/*
  Una manzana representa el un bloque del vecindario, que es lo
  mismo que un espacio en el mapa.
  Se inicializa con su id únicamente y se marca como "a salvo".
*/
func initManzana(id int) Manzana{
  manzana := Manzana{}
  manzana.Id = id
  manzana.Estado = 0
  return manzana
}

/*
  Se inicializan todas las manzanas con su respectiva lista de vecinos
  y se agrega el arreglo de manzanas al vecindario.
*/
func (vecindario *Vecindario) initManzanas(){
  var manzanas []Manzana = make ([]Manzana, NumVertices + 1)
  fmt.Println("numV", NumVertices)
  grafica.BeginTransaction()
  for i := 1; i < NumVertices; i++{
    vecinos := ""
    sonVecinos := 0
    manzanas[i] = initManzana(i)
    for j := 1; j < NumVertices; j++{
      sonVecinos = grafica.GetValue("grafica", i, j)
      if(sonVecinos == 1){
        vecinos += strconv.Itoa(j) + ", "
      }
    }
    query := fmt.Sprintf("INSERT INTO manzanas (ID, VECINOS) VALUES (%d, 0, %s);", i, vecinos)
    grafica.GraphDB.Exec(query)
  }
  grafica.EndTransaction()
  vecindario.Manzanas = manzanas
}

/*
  Cambia el estado de una manzana.
*/
func (manzana *Manzana)SetEstado(estado int){
  manzana.Estado = estado;
}

/*
  mmmh, do I really need this?
*/
func (vecindario *Vecindario) InitFuegoEspecifico(id int){
  fmt.Println("len en init fuego: ", len(vecindario.Manzanas))
  vecindario.Manzanas[id].SetEstado(2)
}

/*
  Se propaga el fuego a todos los vecinos no protegidos de incendios.
*/
func (vecindario *Vecindario) PropagaFuego(){
  incendiados := vecindario.GetIncendiados()

  for i := 0; i < len(incendiados); i++{
    _, vecinos := vecindario.getVecinos(incendiados[i])
    for j := 0; j < len(vecinos); j++{
      manzana := vecindario.Manzanas[vecinos[j]]
      if(manzana.Estado == 0){
        vecindario.Manzanas[vecinos[j]].SetEstado(2)
      }
    }
  }
}

/*
Devuelve un arreglo con el id de las manzanas a salvo
*/
func (vecindario *Vecindario) GetASalvo() []int{
  return vecindario.ConsultaPorEstado(0)
}

/*
Devuelve un arreglo con el id de las manzanas defendidas
*/
func (vecindario *Vecindario) GetDefendidos() []int{
  return vecindario.ConsultaPorEstado(1)
}

/*
  Devuelve un arreglo con el id de las manzanas incendiadas
*/
func (vecindario *Vecindario) GetIncendiados() []int{
  return vecindario.ConsultaPorEstado(2)
}

/*
  Devuelve un arreglo con el id de las manzanas vecinas de incendiados
  que no han sido defendidas ni incendiadas
*/
func (vecindario *Vecindario) GetPorQuemar() []int{
  incendiados := vecindario.GetIncendiados()
  candidatos := []int{}
  for i := 0; i < len(incendiados); i++{
    _, vecinos := vecindario.getVecinos(incendiados[i])
    for j := 0; j < len(vecinos); j++{
      m := vecindario.Manzanas[vecinos[j]]
      if(m.Estado == 0 && !util.Contiene(candidatos, vecinos[j])){
        candidatos = append(candidatos, vecinos[j])
      }
    }
  }
  fmt.Println("")
  return candidatos
}


/*
  Devuelve una copia del vecindario sobre el que se aplica.
*/
func(vecindario *Vecindario) Copia() Vecindario{
  copia := Vecindario{}
  copia.Grado = vecindario.Grado
  copia.Manzanas = vecindario.Manzanas
  return copia
}

/*
  excepción lol
*/
func check(e error){
  if e != nil{
    panic(e)
  }
}

/*
  Evalúa "qué tan bueno" es un escenario tomando en cuenta una relación
  entre los incendios y el número de bomberos utilizados.
*/
func (vecindario *Vecindario) Evalua(numBomberos int) float64{
  quemados := float64(len(vecindario.GetIncendiados()))
  defendidos := float64(len(vecindario.GetDefendidos()))
  return (quemados / float64(vecindario.Cantidad)) * (defendidos / float64(numBomberos))
}

func (vecindario *Vecindario) ConsultaPorEstado(estado int) []int{
  res := []int{}
  for i := 1; i < NumVertices; i++{
    if(vecindario.Manzanas[i].Estado == estado){
      res = append(res, i)
    }
  }
  return res
}

func (vecindario *Vecindario)ConsultaEstado(id int) int {
  return vecindario.Manzanas[id].Estado
}
func consultaVecinos(id int) []int{
  query := fmt.Sprintf("SELECT VECINOS FROM manzanas WHERE ID = %d;", id)
  result, _ := grafica.GraphDB.Query(query)
  var value string
  defer result.Close()
  result.Next()

  values := strings.Split(value, ", ")
  var vecinos []int
  var number int
  for i, _ := range values{
    number, _ = strconv.Atoi(values[i])
    vecinos = append(vecinos, number)
  }
  return vecinos
}

func (vecindario *Vecindario) getVecinos(i int) (string, []int){
  vecinosString := ""
  vecinosArray := []int{}
  grafica.BeginTransaction()
  for j := 1; j < NumVertices; j++ {
    query := fmt.Sprintf("SELECT `%d` FROM grafica WHERE ID = %d;", i, j)
    result, err := grafica.GraphDB.Query(query)
    check(err)
    var value int
    defer result.Close()
    result.Next()
    err = result.Scan(&value)
    check(err)
    if value == 1 {
      vecinosString += strconv.Itoa(j) + ","
      vecinosArray = append(vecinosArray, j)
    }
  }
  // grafica.EndTransaction()
  return vecinosString, vecinosArray
}
