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


/*
Genera la gráfica simple y la carga a la base de datos,
esta gráfica será la matriz de distancias
*/
func GeneraBaseCuadricula(numDiagonales int) int {
  creaBase("grafica")
  beginTransaction()
  // creaBase("recorridos")
  for i := 1; i <= numVertices; i++ {
    for j := 1; j <= numVertices; j++ {
      if (j > numColumnas) && (j % numColumnas == 1) && (j != i){
        fmt.Printf("UNO: %d,%d\n", j, i)
        addRelation("grafica", (j - numColumnas), j, 1)
      } else {
        if (((j - numColumnas) == i) && (numVertices % j != 1)){
          // fmt.Printf("UNO: j - 1 = %d, i = %d \n", j-1, i)
          fmt.Printf("UNO: %d,%d\n", j, i)
          addRelation("grafica", i, j, 1)
        }else{
          if  (((j - 1) == i) || ((j + 1) == i)) && (j % numColumnas != 0){
            fmt.Printf("UNO no modulo: %d, %d\n", j, i)
            addRelation("grafica", i, j, 1)
          }else{
            current := getValue("grafica", i, j)
            fmt.Printf("Vale at %d, %d: %d", j, i, current)
            if getValue("grafica", i, j) == 0 {
              fmt.Printf("INFINITO: %d, %d\n", j, i)
              addRelation("grafica", i, j, 2147483647)
            }
          }
        }
      }
    }
  }
  for i := 1; i <= numVertices; i++ {
    addRelation("grafica", i, i, 0)
  }
  DiagonalesRandom(numDiagonales)
  endTransaction()
  return 1
}


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
      addRelation("grafica", i, j, 1)
      c++
    }
  }
}

func FloydWarshal() (Grafica, Grafica){
  creaBase("recorridos")
  beginTransaction()
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
        distIJ := getValue("grafica", i, j)
        distIK := getValue("grafica", i, k)
        distKJ := getValue("grafica", k, j)
        pathKJ := getValue("recorridos", k, j)
        if(distIJ > distIK + distKJ) {
          addRelation("grafica", i, j, distIK + distKJ)
          addRelation("recorridos", i, j, pathKJ)
        }
      }
    }
  }
  endTransaction()
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
  result, err := GraphDB.Query(query)
  check(err)
  var value int
  defer result.Close()
  result.Next()
  err = result.Scan(&value)

  // intResult, err := result.RowsAffected()
  // check(err)
  return int(value)
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

func beginTransaction()  {
  queryStart := fmt.Sprintf("BEGIN TRANSACTION;")
  _, err := GraphDB.Exec(queryStart)
  check(err)
}

func endTransaction() {
  queryEnd := fmt.Sprintf("END TRANSACTION;")
  _, err := GraphDB.Exec(queryEnd)
  check(err)
}
