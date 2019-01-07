package grafica

import (
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

/*
  Representacion de la grafica
*/
type Grafica struct{
  Nodos [][]int
}

/*
  Genera la gráfica simple que es un arreglo bidimencional, cuando
  el valor es 1, significa que los vértices que representan los índices
  del arreglo son vecinos.
  Esta es la matriz de distancias.
*/
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
  for i := 0; i < len(matriz); i++{
    for j := 0; j < len(matriz); j++{
      if(matriz[i][j] == 0){
        matriz[i][j] = 2147483647
        matriz[j][i] = 2147483647
      }
    }
  }
  grafica := Grafica{}
  grafica.Nodos = matriz
  grafica = grafica.DiagonalesRandom(numDiagonales)
  for i := 0; i < len(matriz); i++{
    matriz[i][i] = 0
  }

  // fmt.Println(grafica)
  return grafica
}

/*
  Agrega diagonales aleatoriamente a una cuadrícula.
*/
func (grafica Grafica) DiagonalesRandom(numDiagonales int) Grafica{
  g := grafica
  g.Nodos = grafica.Nodos
  i := 0
  j := 0
  c := 1
  for c != numDiagonales{
    i = util.RandInt(0, numVertices - numColumnas)
    if((i % numColumnas) == (numColumnas - 1)){
      j = i
    }else{
      j = i + numColumnas + 1
    }
    g.Nodos[i][j] = 1
    g.Nodos[j][i] = 1
    c++
  }
  return g
}

func (grafica Grafica) FloydWarshal() (Grafica, Grafica){
  distancias := grafica.Nodos
  var pats [][]int = make([][]int, numVertices)
  for k := range pats{
    pats[k] = make([]int, numVertices)
  }
  for i := range pats{
    for j := range pats{
      if (distancias[i][j] == 2147483647){
        pats[i][j] = 0
      }else{
        pats[i][j] = i
      }
    }
  }
  for k := range pats{
    for i := range pats{
      for j := range pats{
        if(distancias[i][j] > distancias[i][k] + distancias[k][j]){
          distancias[i][j] = distancias[i][k] + distancias[k][j]
          pats[i][j] = pats[k][j]
        }
      }
    }
  }
  dist := Grafica{}
  dist.Nodos = distancias
  recorridos := Grafica{}
  recorridos.Nodos = pats
  return dist, recorridos
}

/*
  Imprime en texto simplemente la matriz de la gráfica.
*/
func (grafica *Grafica) ImprimeGrafica(nombre string){
  s := ""
  for i := 0; i < len(grafica.Nodos); i++{
    for j := 0; j < len(grafica.Nodos); j++{
      s += strconv.Itoa(grafica.Nodos[i][j]) + ","
    }
    s = s[:len(s) - 1]
    s += "\n"
  }
  s = s[:len(s) - 1]
  f, err := os.Create(nombre)
  check(err)
  defer f.Close()
  n3, err := f.WriteString(s)
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

func check(e error) {
    if e != nil {
        panic(e)
    }
}
