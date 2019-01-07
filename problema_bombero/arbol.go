package problema_bombero

import(
  "container/list"
)

/*
  Vertice, sóĺo guarda indices de padre y hermanos
*/
type Vertice struct{
  Id int
  Padre int
  Hijos []int
}

/*
  Arbol, guarda un arreglo con las referencias de padres e hijos de los vertices
*/
type Arbol struct{
  Elementos int
  Raiz *Vertice
  Vertices []*Vertice
}

/*
  Crea un nuevo vértice (sin padre)
*/
func newVertice(id int) *Vertice{
  ver := Vertice{}
  ver.Id = id
  ver.Hijos = nil
  return &ver
}

/*
  Inicializa nuevo arbol sin referencia a "padre",
  uso numero simbolico de infinito para decir que no tiene padre
*/
func newArbol(r int) *Arbol{
  arbol := Arbol{}
  arbol.Elementos = 0
  arbol.Raiz = newVertice(r)
  arbol.Raiz.Padre = 2147483647
  return &arbol
}

/*
  Crea árbol completo a partir de un vértice dado de la gráfica.
*/
func CreaArbol(distancias [][]int, raiz int, numVertices int) *Arbol{
  arbol := newArbol(raiz)
  porProcesar := list.New()
  porProcesar.PushFront(raiz)
  vertices := make([]*Vertice, numVertices)
  arbol.Vertices = vertices
  arbol.Vertices[arbol.Raiz.Id] = arbol.Raiz
  arbol.Elementos++

  return arbol.agregaHijos(raiz, distancias, numVertices, porProcesar, arbol.Vertices)
}

/*
  Hace el trabajo sucio de agregar hijos del arbol de manera recursiva.
*/
func (arbol *Arbol) agregaHijos(id int, dist [][]int, numVer int, porProc *list.List, vert []*Vertice) *Arbol{
  var verTemp *Vertice
  for i := 0; i < len(dist); i++ {
      if(dist[id][i] == 1 && vert[i] == nil){
          verTemp = newVertice(i)
          porProc.PushBack(verTemp.Id)
          arbol.Vertices[id].Hijos = append(arbol.Vertices[id].Hijos, verTemp.Id)
          verTemp.Padre = id
          vert[verTemp.Id] = verTemp
          arbol.Elementos++
      }
  }
  porProc.Remove(porProc.Front())
  if(porProc.Len() == 0){
    return arbol
  }else{
    temp := porProc.Front().Value

    return arbol.agregaHijos(temp.(int), dist, numVer, porProc, vert)
  }
}

/*
  Obtiene un arreglo con el camino más corto entre un vértice y el vértice S de su árbol
*/
func (arbol *Arbol) GetTrayectoria(v int) []int{
  var result []int
  result = append(result, v)
  var pa = arbol.Vertices[v].Padre
  for pa != 2147483647{
    result = append(result, pa)
    pa = arbol.Vertices[pa].Padre
  }
  return result
}
