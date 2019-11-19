package util

import(
  "container/list"
  // "fmt"
)

/*
  VerticeD, guarda referencias a padres e hijos
*/
type VerticeD struct{
  Id int
  Padres []int
  Hijos []int
  Nivel int
}

/*
  Grafica dirigida, guarda en un arreglo las referencias a padres e hijos de cada vertice
*/
type Dirigida struct{
  Elementos int
  Raices []*VerticeD
  Vertices []*VerticeD
}

/*
  Crea un nuevo vertice
*/
func newVerticeD(id int, niv int) *VerticeD{
  ver := VerticeD{}
  ver.Id = id
  ver.Padres = nil
  ver.Hijos = nil
  ver.Nivel = niv
  return &ver
}

/*
  Inicializa una grafica dirigida
*/
func newDirigida(r []int) *Dirigida{
  dirigida := Dirigida{}
  dirigida.Elementos = 0
  // aux := make([]*VerticeD, len(r))
  vAux := newVerticeD(0, 0)
  for _, raiz := range r{
    vAux = newVerticeD(raiz, 0)
    vAux.Padres = []int{2147483647}
    dirigida.Raices = append(dirigida.Raices, vAux)
  }
  return &dirigida
}

/*
  Crea grafica dirigida para la instancia
*/
func CreaDirigida(distancias [][]int, raices []int, numVertices int) *Dirigida{
  dirigida := newDirigida(raices)
  porProcesar := list.New()
  vertices := make([]*VerticeD, numVertices)
  dirigida.Vertices = vertices
  for _, r := range dirigida.Raices{
    porProcesar.PushFront(r.Id)
    dirigida.Vertices[r.Id] = r
    dirigida.Elementos++
  }
  temp := porProcesar.Front().Value
  return dirigida.agregaVertices(temp.(int), distancias, porProcesar, dirigida.Vertices)
}

/*
  Trabajo sucio de agregar los vertices
*/
func (dirigida *Dirigida) agregaVertices(id int, dist [][]int, porProc *list.List, vert []*VerticeD) *Dirigida{
  var verTemp *VerticeD
  for i := 0; i < len(dist); i++ {
    if(dist[id][i] == 1 && (vert[i] == nil || vert[i].Nivel < vert[id].Nivel)){
      if(vert[i] == nil){
        verTemp = newVerticeD(i, (vert[id].Nivel + 1))
        vert[i] = verTemp
        porProc.PushBack(i)
      }
      if(dirigida.Vertices[id].Hijos != nil){
        dirigida.Vertices[id].Hijos = append(dirigida.Vertices[id].Hijos, verTemp.Id)
      }
      vert[i].Padres = append(vert[i].Padres, id)
      dirigida.Elementos++
    }
  }
  porProc.Remove(porProc.Front())
  if(porProc.Len() == 0){
    return dirigida
  }else{
    temp := porProc.Front().Value
    return dirigida.agregaVertices(temp.(int), dist, porProc, vert)
  }
}
