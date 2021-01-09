package problema_bombero

 import(
  "github.com/kaarla/HOC-Proyecto2/util"
  "github.com/kaarla/HOC-Proyecto2/problema_bombero/grafica"
  "fmt"
  // "container/list"
  // "strings"
  // "io/ioutil"
  "strconv"
  "github.com/fatih/set"
 )

 //arreglo de trayectorias
var Trayectorias [][]int
//arreglo de distancias
//var Distancias [][]int


 type Vecindario struct{
   Manzanas []Manzana
   // Mapa [][]int
   Grado int
 }

type Manzana struct{
  Id int
  Estado int
  Vecinos []int
}


func NewVecindario(mapa [][]int) *Vecindario{
  vec := Vecindario{}
  // vec.Mapa = mapa
  return &vec
}

func VecindarioCero(grafica string) Vecindario{
  vecindario := Vecindario{}
  // vecindario.Mapa = InitMapa(grafica)
  vecindario.Manzanas = nil
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
  manzana.Vecinos = nil
  return manzana
}

/*
  Se inicializan todas las manzanas con su respectiva lista de vecinos
  y se agrega el arreglo de manzanas al vecindario.
*/
func (vecindario *Vecindario) initManzanas(){
  grafica.BeginTransaction()
  for i := 0; i < NumVertices; i++{
    vecinos := ""
    sonVecinos := 0
    for j := 0; j < NumVertices; j++{
      sonVecinos = grafica.GetValue("grafica", i, j)
      if(sonVecinos == 1){
        vecinos += strconv.Itoa(j) + ", " //append(vecinos, j)
      }
    }
    query := fmt.Sprintf("INSERT INTO manzanas (ID, ESTADO, VECINOS) VALUES (%d, 0, %s);", i, vecinos)
    grafica.GraphDB.Exec(query)
    // vecindario.Manzanas = manzanas
  }
  grafica.EndTransaction()
}

/*
  Cambia el estado de una manzana.
*/
func (manzana *Manzana) SetEstado(estado int, id int){
  grafica.BeginTransaction()
  query := fmt.Sprintf("UPDATE grafica SET ESTADO = %d WHERE ID = %d;", estado, id)
  grafica.GraphDB.Exec(query)
  grafica.EndTransaction()
}

/*
  mmmh, do I really need this?
*/
func (vecindario *Vecindario) InitFuegoEspecifico(manzana int){
    vecindario.Manzanas[manzana].Estado = 2
}

/*
  Se propaga el fuego a todos los vecinos no protegidos de incendios.
*/
func (vecindario *Vecindario) PropagaFuego(){
  incendiados := vecindario.GetIncendiados()
  // defendidos := vecindario.GetDefendidos()
  // asalvo := vecindario.GetASalvo()
  // // fmt.Println("<p> incenciados", incendiados, "</p>")
  // // fmt.Println("<p> defendidos", defendidos, "</p>")
  // // fmt.Println("<p> asalvo", asalvo, "</p>")
  // // fmt.Println("<p>vecinos de 3, ", vecindario.Manzanas[3].Vecinos, "</p>")
  for i := 0; i < len(incendiados); i++{
    v := vecindario.Manzanas[incendiados[i]].Vecinos
    for j := 0; j < len(v); j++{
      m := vecindario.Manzanas[v[j]]
      if(m.Estado == 0){
        vecindario.Manzanas[v[j]].Estado = 2
      }
    }
  }
}

/*
Devuelve un arreglo con el id de las manzanas a salvo
*/
func (vecindario *Vecindario) GetASalvo() []int{
  return consultaPorEstado(0)
}

/*
Devuelve un arreglo con el id de las manzanas defendidas
*/
func (vecindario *Vecindario) GetDefendidos() []int{
  return consultaPorEstado(1)
}

/*
  Devuelve un arreglo con el id de las manzanas incendiadas
*/
func (vecindario *Vecindario) GetIncendiados() []int{
  return consultaPorEstado(2)
}

/*
  Devuelve un arreglo con el id de las manzanas vecinas de incendiados
  que no han sido defendidas ni incendiadas
*/
func (vecindario *Vecindario) GetPorQuemar() []interface{}{
  incendiados := vecindario.GetIncendiados()
  conjuntoCandidatos := set.New(set.ThreadSafe)
  for _, id := range incendiados {
    vecinos := getVecinos(id)
    query := fmt.Sprintf("SELECT ID FROM manzanas WHERE ID IN (%s) AND ESTADO = 0;", vecinos[:len(vecinos)-1])
    result, err := grafica.GraphDB.Query(query)
    check(err)
    var value int
    defer result.Close()
    for result.Next() {
      err = result.Scan(&value)
      conjuntoCandidatos.Add(value)
    }
  }
  return conjuntoCandidatos.List()
}


/*
  Devuelve una copia del vecindario sobre el que se aplica.
*/
func(vecindario *Vecindario) Copia() Vecindario{
  copia := Vecindario{}
  // copia.Mapa = vecindario.Mapa
  copia.Manzanas = vecindario.Manzanas
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
  return (quemados / float64(len(vecindario.Manzanas))) * (defendidos / float64(numBomberos))
}

func consultaPorEstado(estado int) []int{
  query := fmt.Sprintf("SELECT ID FROM manzanas WHERE ESTADO = %d;", estado)
  result, err := grafica.GraphDB.Query(query)
  var indices []int
  check(err)
  var value int
  defer result.Close()
  for result.Next() {
    err = result.Scan(&value)
    indices = append(indices, value)
  }
  return indices
}

func consultaEstado(id int) int {
  query := fmt.Sprintf("SELECT ESTADO FROM manzanas WHERE ID = %d;", id)
  result, err := grafica.GraphDB.Query(query)
  check(err)
  var value int
  defer result.Close()
  result.Next()
  err = result.Scan(&value)

  return int(value)
}


/*
  Formato para imprimir una manzana con su color con javascript.
*/
func (vecindario *Vecindario) PrintManzana(){
  color := ""
  for _, m := range vecindario.Manzanas{
    switch m.Estado {
    case 0:
      color = "pink}"
    case 1:
      color = "blue}"
    case 2:
      color = "red}"
    }
    if(util.Contiene(PorSalvar, m.Id)){
      color = "green}"
    }
     fmt.Println(m.Id, " {color:", color)
  }
}

func getVecinos(i int) string{
  vecinos := ""
  grafica.BeginTransaction()
  result := ""
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
      vecinos += strconv.Itoa(j) + ","
    }
  }
  grafica.EndTransaction()
  return result
}

/*
  Para dar formato de SVG.
*/
func (vecindario *Vecindario) PrintSVG(){
    x := 5
    y := 5
    numColumnas := 3
    h := 300
    color := ""
    switch len(vecindario.Manzanas) {
    case 9:
      numColumnas = 3
      h = 250
    case 50:
      numColumnas = 5
      h =600
    case 100:
      numColumnas = 10
      h = 600
    case 1000:
      numColumnas = 40
      h = 1500
    }
    fmt.Printf("<svg height=\"%d\" width=\"2000\">\n<g font-size=\"10\" font-family=\"sans-serif\" fill=\"black\" stroke=\"none\">\n", h)
    // fmt.Println("------------------------", len(vecindario.Manzanas))
    for _, m := range vecindario.Manzanas{
      switch m.Estado {
      case 0:
        color = "pink"
      case 1:
        color = "blue"
      case 2:
        color = "red"
      }
      if(util.Contiene(PorSalvar, m.Id)){
        color = "green"
      }
      // fmt.Println("m.Id", "numColumnas", m.Id, numColumnas)
      if(m.Id % numColumnas == 0){
        x = 5
        y += 50
      }
      fmt.Printf("<circle id=\"point%d\" cx=\"%d\" cy=\"%d\" r=\"6\" fill=\"%s\" stroke=\"%s\" />\n",
        m.Id, x, y, color, color)
      fmt.Printf("<text x=\"%d\" y=\"%d\" dy=\"%d\">%d</text>\n", x, y, -10, m.Id)
      for _, n := range m.Vecinos{
        switch n {
        case m.Id + 1:
          fmt.Printf("<path id=\"line%d%d\" d=\"M %d %d l %d %d\" stroke=\"black\" stroke-width=\"3\" />\n",
             m.Id, n, x, y, 47, 0)
        case m.Id + numColumnas:
          fmt.Printf("<path id=\"line%d%d\" d=\"M %d %d l %d %d\" stroke=\"black\" stroke-width=\"3\" />\n",
             m.Id, n, x, y, 0, 47)
        case m.Id + numColumnas + 1:
          fmt.Printf("<path id=\"line%d%d\" d=\"M %d %d l %d %d\" stroke=\"black\" stroke-width=\"3\" />\n",
             m.Id, n, x, y, 47, 47)
        }
      }
      x += 50
    }
  fmt.Println("</g>\n</svg>")
}
