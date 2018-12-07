package problema_bombero

import (
  "fmt"
  // "strconv"
  // "math/rand"
  // "io/ioutil"
  // "strings"
  // "bytes"
  "container/list"
)

type Vertice struct{
  Id int
  // Estado int
  Padre int
  Hijos []int
}

type Arbol struct{
  Elementos int
  Raiz *Vertice
  Vertices []*Vertice
}

func newVertice(id int) *Vertice{
  ver := Vertice{}
  ver.Id = id
  ver.Hijos = nil
  return &ver
}

func newArbol(r int) *Arbol{
  arbol := Arbol{}
  arbol.Elementos = 0
  arbol.Raiz = newVertice(r)
  arbol.Raiz.Padre = 2147483647
  return &arbol
}

func CreaArbol(distancias [][]int, raiz int, numVertices int) *Arbol{
  arbol := newArbol(raiz)
  fmt.Println("raiz padre", arbol.Raiz.Padre)
  porProcesar := list.New()
  porProcesar.PushFront(raiz)
  var verTemp *Vertice = newVertice(raiz)
  vertices := make([]*Vertice, numVertices)
  arbol.Vertices = vertices
  arbol.Vertices[verTemp.Id] = verTemp
  arbol.Elementos++
  return arbol.agregaHijos(raiz, distancias, numVertices, porProcesar, arbol.Vertices)
}

func (arbol *Arbol) agregaHijos(id int, dist [][]int, numVer int, porProc *list.List, vert []*Vertice) *Arbol{
  if(porProc.Len() == 0){
    return arbol
  }
  var verTemp *Vertice
  for i := 0; i < len(dist); i++ {
      if(dist[id][i] == 1 && vert[i] == nil){
          verTemp = newVertice(i)
          porProc.PushBack(verTemp.Id)
          arbol.Raiz.Hijos = append(arbol.Raiz.Hijos, verTemp.Id)
          if(vert[i] == nil){
            // fmt.Println("agregando ", i)
            verTemp.Padre = id
            vert[verTemp.Id] = verTemp
            arbol.Elementos++
          }
      }
  }
  porProc.Remove(porProc.Front())
  if(porProc.Len() == 0){
    return arbol
  }else{
  temp := porProc.Front().Value
  return arbol.agregaHijos(temp.(int), dist, numVer, porProc, vert)}
}
