package problema_bombero

 import(
  "github.com/kaarla/HOC-Proyecto2/problema_bombero/grafica"
  "fmt"
  "github.com/kaarla/HOC-Proyecto2/util"
  "strings"
  "strconv"
  // "github.com/fatih/set"
 )

 //arreglo de trayectorias
var Trayectorias [][]int


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
  var manzanas []Manzana = make ([]Manzana, NumVertices)
  fmt.Println("numV", NumVertices)
  grafica.BeginTransaction()
  for i := 0; i < NumVertices; i++{
    vecinos := ""
    sonVecinos := 0
    manzanas[i] = initManzana(i)
    for j := 0; j < NumVertices; j++{
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
  fmt.Println("len de manzanas en init", len(vecindario.Manzanas))
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
  for i := 0; i < len(vecindario.Manzanas); i++{
    if(vecindario.Manzanas[i].Estado == estado){
      res = append(res, i)
    }
  }
  return res
}

func (vecindario *Vecindario)ConsultaEstado(id int) int {
  fmt.Println("len de manzanas", len(vecindario.Manzanas))
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


/*
  Formato para imprimir una manzana con su color con javascript.
*/
// func (vecindario *Vecindario) PrintManzana(){
//   color := ""
//   for _, m := range vecindario.Manzanas{
//     switch m.Estado {
//     case 0:
//       color = "pink}"
//     case 1:
//       color = "blue}"
//     case 2:
//       color = "red}"
//     }
//     if(util.Contiene(PorSalvar, m.Id)){
//       color = "green}"
//     }
//      fmt.Println(m.Id, " {color:", color)
//   }
// }

func (vecindario *Vecindario) getVecinos(i int) (string, []int){
  vecinosString := ""
  vecinosArray := []int{}
  grafica.BeginTransaction()
  for j := 1; j <= NumVertices; j++ {
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
  grafica.EndTransaction()
  return vecinosString, vecinosArray
}

/*
  Para dar formato de SVG.
*/
// func (vecindario *Vecindario) PrintSVG(){
//     x := 5
//     y := 5
//     numColumnas := 3
//     h := 300
//     color := ""
//     switch len(vecindario.Manzanas) {
//     case 9:
//       numColumnas = 3
//       h = 250
//     case 50:
//       numColumnas = 5
//       h =600
//     case 100:
//       numColumnas = 10
//       h = 600
//     case 1000:
//       numColumnas = 40
//       h = 1500
//     }
//     fmt.Printf("<svg height=\"%d\" width=\"2000\">\n<g font-size=\"10\" font-family=\"sans-serif\" fill=\"black\" stroke=\"none\">\n", h)
//     // fmt.Println("------------------------", len(vecindario.Manzanas))
//     for _, m := range vecindario.Manzanas{
//       switch m.Estado {
//       case 0:
//         color = "pink"
//       case 1:
//         color = "blue"
//       case 2:
//         color = "red"
//       }
//       if(util.Contiene(PorSalvar, m.Id)){
//         color = "green"
//       }
//       // fmt.Println("m.Id", "numColumnas", m.Id, numColumnas)
//       if(m.Id % numColumnas == 0){
//         x = 5
//         y += 50
//       }
//       fmt.Printf("<circle id=\"point%d\" cx=\"%d\" cy=\"%d\" r=\"6\" fill=\"%s\" stroke=\"%s\" />\n",
//         m.Id, x, y, color, color)
//       fmt.Printf("<text x=\"%d\" y=\"%d\" dy=\"%d\">%d</text>\n", x, y, -10, m.Id)
//       for _, n := range m.Vecinos{
//         switch n {
//         case m.Id + 1:
//           fmt.Printf("<path id=\"line%d%d\" d=\"M %d %d l %d %d\" stroke=\"black\" stroke-width=\"3\" />\n",
//              m.Id, n, x, y, 47, 0)
//         case m.Id + numColumnas:
//           fmt.Printf("<path id=\"line%d%d\" d=\"M %d %d l %d %d\" stroke=\"black\" stroke-width=\"3\" />\n",
//              m.Id, n, x, y, 0, 47)
//         case m.Id + numColumnas + 1:
//           fmt.Printf("<path id=\"line%d%d\" d=\"M %d %d l %d %d\" stroke=\"black\" stroke-width=\"3\" />\n",
//              m.Id, n, x, y, 47, 47)
//         }
//       }
//       x += 50
//     }
//   fmt.Println("</g>\n</svg>")
// }
