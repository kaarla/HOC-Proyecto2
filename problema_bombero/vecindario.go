package problema_bombero

 import(
  "github.com/kaarla/HOC-Proyecto2/util"
  "fmt"
  "strings"
  "io/ioutil"
  "strconv"
 )

type Manzana struct{
  Id int
  Estado int
  Vecinos []int
}

type Vecindario struct{
  Manzanas []Manzana
  Mapa [][]int
  Grado int
}


func NewVecindario(mapa [][]int) *Vecindario{
  vec := Vecindario{}
  vec.Mapa = mapa
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

/*
Inicializa el mapa del vecindario a partir de la gráfica
que recibe en formato de texto convirtiéndolo en un arreglo bidimensional.
*/
func initMapa(grafica string) [][]int{
  datos, err := ioutil.ReadFile(grafica)
  check(err)
  lineas := strings.Split(string(datos), "\n") //cada línea es una fila de la cuadrícula

  var mapa [][]int = make([][]int, len(lineas) - 1)
  for k := range mapa{
    mapa[k] = make([]int, len(lineas) - 1)
  }
  for i := 0; i < len(lineas) - 1; i++{
    linea := strings.Split(string(lineas[i]), ",")
    for j := 0; j < len(lineas) - 1; j++{
      num, err := strconv.ParseInt(linea[j], 10, 64)
      check(err)
      mapa[i][j] = int(num)
      mapa[j][i] = int(num)
    }
  }
  return mapa
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

/*
  Cambia el estado de una manzana.
*/
func (manzana *Manzana) SetEstado(estado int){
  manzana.Estado = estado
}

/*
  mmmh, do I really need this?
*/
func (vecindario *Vecindario) InitFuegoEspecifico(manzana int){
    vecindario.Manzanas[manzana].SetEstado(2)
}

/*
  Se propaga el fuego a todos los vecinos no protegidos de incendios.
*/
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

/*
  Devuelve un arreglo con el id de las manzanas incendiadas
*/
func (vecindario *Vecindario) GetIncendiados() []int{
  res := []int{}
  for i := 0; i < len(vecindario.Manzanas); i++{
    if(vecindario.Manzanas[i].Estado == 2){
      res = append(res, i)
    }
  }
  return res
}

/*
  Devuelve un arreglo con el id de las manzanas a salvo
*/
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

/*
  Devuelve un arreglo con el id de las manzanas defendidas
*/
func (vecindario *Vecindario) GetDefendidos() []int{
  res := []int{}
  for i := 0; i < len(vecindario.Manzanas); i++{
    if(vecindario.Manzanas[i].Estado == 1){
      res = append(res, i)
    }
  }
  return res
}

/*
  Devuelve un arreglo con el id de las manzanas vecinas de incendiados
  que no han sido defendidas ni incendiadas
*/
func (vecindario *Vecindario) GetCandidatos() []int{
  incendiados := vecindario.GetIncendiados()
  candidatos := []int{}
  for i := 0; i < len(incendiados); i++{
    v := vecindario.Manzanas[incendiados[i]].Vecinos
    for j := 0; j < len(v); j++{
      m := vecindario.Manzanas[v[j]]
      if(m.Estado == 0 && !util.Contiene(candidatos, v[j])){
        candidatos = append(candidatos, v[j])
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
  copia.Mapa = vecindario.Mapa
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
     fmt.Println(m.Id, " {color:", color)
  }
}

/*
  Para dar formato de SVG.
*/
func (vecindario *Vecindario) PrintSVG(){
    x := 5
    y := 5
    numColumnas := 0
    h := 0
    color := ""
    switch len(vecindario.Manzanas) {
    case 9:
      numColumnas = 3
      h = 500
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
    for _, m := range vecindario.Manzanas{
      switch m.Estado {
      case 0:
        color = "pink"
      case 1:
        color = "blue"
      case 2:
        color = "red"
      }
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
