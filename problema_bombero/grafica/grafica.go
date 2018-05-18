package grafica

import (
  "fmt"
  "strconv"
  "math/rand"
  // "strings"
)

var(
  numVertices = 50
  numfilas = 10
  numColumnas = 5
)

type Grafica struct{
  nodos [][]int
}

func GeneraCuadricula(numDiagonales int) Grafica{
  var matriz [][]int = make([][]int, numVertices)
  for k := range matriz{
    matriz[k] = make([]int, numVertices)
  }
  for i := 0; i < len(matriz); i++{
    for j := 0; j < len(matriz); j++{
      if((j >= numColumnas) && (j % numColumnas == 0)){
        matriz[j - numColumnas][j] = 1
        matriz[j][j - numColumnas] = 1
      }else{
        if((j - numColumnas) == i && numVertices % j != 0){
          matriz[i][j] = 1
          matriz[j][i] = 1
        }
        if((j - 1) == i){
          matriz[i][j] = 1
          matriz[j][i] = 1
        }
      }
    }
  }
  grafica := Grafica{}
  grafica.nodos = matriz
  grafica = grafica.DiagonalesRandom(numDiagonales)
  // fmt.Println(grafica)
  return grafica
}

func (grafica Grafica) DiagonalesRandom(numDiagonales int) Grafica{
  g := grafica
  g.nodos = grafica.nodos
  i := 0
  j := 0
  c := 1
  for c != numDiagonales{
    i = randInt(0, numVertices - numColumnas)
    if((i % numColumnas) == (numColumnas - 1)){
      j = i
    }else{
      j = i + numColumnas + 1
    }
    g.nodos[i][j] = 1
    g.nodos[j][i] = 1
    c++
  }
  return g
}

func (grafica *Grafica) ImprimeGrafica(){
  s := ""
  for i := 0; i < len(grafica.nodos); i++{
    for j := 0; j < len(grafica.nodos); j++{
      s += strconv.Itoa(grafica.nodos[i][j]) + ","
    }
    s = s[:len(s) - 1]
    s += "\n"
  }
  s = s[:len(s) - 1]
  fmt.Println(s)
}

func (grafica *Grafica) ImprimeV(){
  s := ""
  for i := 0; i < len(grafica.nodos); i++{
    for j := 0; j < len(grafica.nodos); j++{
      if(grafica.nodos[i][j] > 0){
        s += strconv.Itoa(i) + " -> " + strconv.Itoa(j) + "\n"
      }
    }
  }
  fmt.Println(s)
}

func randInt(min int, max int) int {
  return min + rand.Intn(max-min)
}
