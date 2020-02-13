package grafica

import (
  "database/sql"
  "fmt"
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

func GeneraBaseCuadricula(numDiagonales int) int {
  creaBase("grafica")
  creaBase("recorridos")
  for i := 1; i <= numVertices; i++ {
    for j := 1; j <= numVertices; j++ {
      if (j > numColumnas) && (j % numColumnas == 1) {
        addRelation("grafica", (j - numColumnas), j, 1)
      } else {
        if (((j - numColumnas) == i) && (numVertices % j != 1)) || ((j - 1) == i) {
          addRelation("grafica", i, j, 1)
        }else if j == i {
          addRelation("grafica", i, j, 0)
        }else{
          addRelation("grafica", i, j, 2147483647)
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
  for i := 1; i <= numVertices; i++ {
    for j := 1; j <= numVertices; j++ {
      result := getValue("grafica", j, i)
      if (result == 2147483647){
        addRelation("recorridos", i, j, 0)
      }else{
        addRelation("recorridos", i, j, i)
      }
    }
  }
  for k := 1; k <= numVertices; k++ {
    for i := 1; i <= numVertices; i++ {
      for j := 1; j <= numVertices; j++ {
        distIJ := getValue("grafica", j, i)
        distIK := getValue("grafica", k, i)
        distKJ := getValue("grafica", j, k)
        pathKJ := getValue("recorridos", j, k)
        if(distIJ > distIK + distKJ) {
          addRelation("grafica", i, j, distIK + distKJ)
          addRelation("recorridos", i, j, pathKJ)
        }
      }
    }
  }
  dist := Grafica{}
  recorridos := Grafica{}
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

func addRelation(name string, i int, j int, dist int){
  query := fmt.Sprintf("UPDATE %s SET `%d` = %d WHERE ID = %d;", name, i, dist, j)
  _, err := GraphDB.Exec(query)
  check(err)
  query = fmt.Sprintf("UPDATE %s SET `%d` = %d WHERE ID = %d;", name, j, dist, i)
  _, err = GraphDB.Exec(query)
  check(err)
}

func getValue(name string, i int, j int) int{
  query := fmt.Sprintf("SELECT `%d` FROM %s WHERE ID = %d;", i, name, j)
  result, err := GraphDB.Exec(query)
  check(err)
  intResult, err := result.RowsAffected()
  check(err)
  return int(intResult)
}

func check(e error) {
    if e != nil {
      fmt.Println(e)
      panic(e)
    }
}

func creaBase(name string) {
  var err error
  for i := 1; i <= numVertices; i++ {
    queryAddFila := fmt.Sprintf("INSERT INTO %s (ID) VALUES (%d);", name, i)
    _, err = GraphDB.Exec(queryAddFila)
    check(err)
    queryAddColumna := fmt.Sprintf("ALTER TABLE %s ADD `%d` INT;", name, i)
    _, err = GraphDB.Exec(queryAddColumna)
    check(err)
  }
}
