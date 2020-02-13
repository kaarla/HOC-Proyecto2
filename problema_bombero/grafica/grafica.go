package grafica

import (
  "database/sql"
  // "errors"
  "fmt"
  "strconv"
  "os"
  "github.com/kaarla/HOC-Proyecto2/util"
)

var(
  numVertices = 9
  numfilas = 3
  numColumnas = 3
)

var GraphDB *sql.DB

/*
  Representacion de la grafica
*/
type Grafica struct{
  Nodos [][]int
}

type intNil struct{
  valid bool
  value int
}


func GeneraBaseCuadricula(numDiagonales int) int {
  queryAddColumna := ""
  queryAddFila := ""
  var err error
  for i := 1; i <= numVertices; i++ {
    queryAddFila = fmt.Sprintf("INSERT INTO grafica (ID) VALUES (%d);", i)
    _, err = GraphDB.Exec(queryAddFila)
    check(err)
    queryAddColumna = fmt.Sprintf("ALTER TABLE grafica ADD `%d` INT;", i)
    _, err = GraphDB.Exec(queryAddColumna)
    check(err)
  }
  for i := 1; i <= numVertices; i++ {
    for j := 1; j <= numVertices; j++ {
      if (j > numColumnas) && (j % numColumnas == 1) {
        addRelation("grafica", (j - numColumnas), j, 1)
      } else {
        if (((j - numColumnas) == i) && (numVertices % j != 1)) || ((j - 1) == i) {
          addRelation("grafica", i, j, 1)
        }
      }
    }
  }
  DiagonalesRandom(numDiagonales)
  return 1
}

/*
  Genera la gráfica simple que es un arreglo bidimencional, cuando
  el valor es 1, significa que los vértices que representan los índices
  del arreglo son vecinos.
  Esta es la matriz de distancias.
*/


/*
  Agrega diagonales aleatoriamente a una cuadrícula.
*/
func DiagonalesRandom(numDiagonales int){
  i := 0
  j := 0
  c := 1
  for c != numDiagonales{
    i = util.RandInt(0, numVertices - numColumnas)
    if((i % numColumnas) == 0){
      j = i
    }else{
      j = i + numColumnas + 1
    }
    addRelation("grafica", i, j, 1)
    c++
  }
}

func FloydWarshal() (Grafica, Grafica){
  // queryInicial := ""
  // distancias := grafica.Nodos
  // var paths [][]int = make([][]int, numVertices)
  // for k := range paths{
  //   paths[k] = make([]int, numVertices)
  // }
  for i := 1; i <= numVertices; i++ {
    for j := 1; j <= numVertices; j++ {
      result := getValue("grafica", i, j)
      if (result == 2147483647){
        addRelation("recorridos", i, j, 0)
        // paths[i][j] = 0
      }else{
        addRelation("recorridos", i, j, i)
        // paths[i][j] = i
      }
    }
  }
  for k := 1; k <= numVertices; k++ {
    for i := 1; i <= numVertices; i++ {
      for j := 1; j <= numVertices; j++ {
        distIJ := getValue("grafica", i, j)
        distIK := getValue("grafica", i, k)
        distKJ := getValue("grafica", k, j)
        pathKJ := getValue("recorridos", k, j)
        // if(distancias[i][j] > distancias[i][k] + distancias[k][j]){
        if(distIJ > distIK + distKJ) {
          // distancias[i][j] = distancias[i][k] + distancias[k][j]
          addRelation("grafica", i, j, distIK + distKJ)
          // paths[i][j] = paths[k][j]
          addRelation("recorridos", i, j, pathKJ)
        }
      }
    }
  }
  dist := Grafica{}
  // dist.Nodos = distancias
  recorridos := Grafica{}
  // recorridos.Nodos = paths
  return dist, recorridos
}

/*
  Imprime en texto simplemente la matriz de la gráfica.
*/
func (grafica *Grafica) ImprimeGrafica(nombre string){
  s := ""
  // for i := 0; i < len(grafica.Nodos); i++{
  //   for j := 0; j < len(grafica.Nodos); j++{
  //     s += strconv.Itoa(grafica.Nodos[i][j]) + ","
  //   }
  //   s = s[:len(s) - 1]
  //   s += "\n"
  // }
  // s = s[:len(s) - 1]
  f, _ := os.Create(nombre)
  // check(err)
  // defer f.Close()
  n3, _ := f.WriteString(s)
  fmt.Printf("se escribieron %d bytes en %s\n", n3, nombre)
}

/*
  No necesito esto .-.
*/
func (grafica *Grafica) ImprimeV(){
  s := ""
  for i := 0; i < len(grafica.Nodos); i++{
    for j := 0; j < len(grafica.Nodos); j++{
      if(grafica.Nodos[i][j] == 1){
        s += strconv.Itoa(i) + " -> " + strconv.Itoa(j) + "\n"
      }
    }
  }
  fmt.Println(s)
}

func addRelation(name string, i int, j int, dist int){
  query := fmt.Sprintf("UPDATE %s SET `%d` = %d WHERE ID = %d;", name, i, dist, j)
  _, err := GraphDB.Exec(query)
  check(err)
  query = fmt.Sprintf("UPDATE %s SET `%d` = %d WHERE ID = %d;", name, j, dist, i)
  _, err = GraphDB.Exec(query)
  check(err)
}

func getValue(name string, i int, j int) int {
  query := fmt.Sprintf("SELECT `%d` FROM %s WHERE ID = %d;", i, name, j)
  result, err := GraphDB.Query(query)
  check(err)
  defer result.Close()
  result.Next()
  // var jValue intNil
  // err = result.Scan(&jValue)
  check(err)
  fmt.Printf("result: %s", result)
  // if jValue == nil{
    return 2147483647
  // }

  // return int(jValue)
}

func check(e error) {
    if e != nil {
      fmt.Println(e)
      panic(e)
    }
}
