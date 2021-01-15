package problema_bombero
import(
  "fmt"
  "github.com/kaarla/HOC-Proyecto2/util"
  "github.com/kaarla/HOC-Proyecto2/problema_bombero/grafica"
  "math/rand"
)

//Modela un candidato a defender
type Candidato struct{
  Id int
  NumTrayectorias int
  DistMinB int
  DistMinS int
  Prioridad float64
}

/*
  Crea un nuevo candidato a partir de
*/
func NewCandidato(id int, t int) *Candidato{
  cand := new(Candidato)
  cand.Id = id
  cand.NumTrayectorias = t
  cand.DistMinB = 2147483647
  cand.DistMinS = 2147483647
  cand.Prioridad = 0.0
  return cand
}

/*
  Encuentra el Id de los vertices mas cercanos en fuego y por salvar del candidato
*/
func (c *Candidato) FindMins(dists []int, manzanas []Manzana){
  var minS int
  var minB int
  auxS := false
  auxB := false
  i := 0
  for !(auxB && auxS){
    if(ConsultaEstado(i) == 2 && !auxB){
      minB = grafica.GetValue("distancias", c.Id, i)
      c.DistMinB = minB
      auxB = true
    }
    if(util.Contiene(PorSalvar, manzanas[i].Id) && !auxS){
      minS = grafica.GetValue("distancias", c.Id, i)
      c.DistMinS = minS
      auxS = true
    }
    i++
  }
}

/*
  Compara dos candidatos
  si es mayor, devuelve 1
  si es igual, devuelve 0
  si es menor, devuelve -1
*/
func (cA *Candidato) compareTo(cB *Candidato) int{
  if(cA.DistMinS == 1 && cA.DistMinB == 1){
    fmt.Println("<p>", cA.Id, "mayor que", cB.Id, "</p>")
    return 1
  }
  if(cB.DistMinS == 1 && cB.DistMinB == 1){
    return 1
  }
  //falta incluir prioridad obtenida con BFS
  if(cA.NumTrayectorias == cB.NumTrayectorias){
    if(cA.DistMinS < cB.DistMinS){ //prioridad al que esta mas cerca de S
        return 1
    }else if(cA.DistMinS > cB.DistMinS){
      return -1
    }//son iguales sus distancias a S
    if(cA.DistMinB > cB.DistMinB){
      return 1
    }else if(cA.DistMinB < cB.DistMinB){
      return -1
    }//son iguales sus distancias a B
    return 0
  }else if(cA.NumTrayectorias > cB.NumTrayectorias){
    return 1
  }
  return -1
}

/*
  QuickSort para los candidatos
*/
func QSort(cs []*Candidato) []*Candidato{
  if len(cs) < 2{
    return cs
  }
  izq, der := 0, (len(cs) - 1)
  p := rand.Int() % len(cs)
  cs[p], cs[der] = cs[der], cs[p]
  for i, _ := range cs{
    if cs[i].compareTo(cs[der]) == -1{
      cs[izq], cs[i] = cs[i], cs[izq]
      izq++
    }
  }
  cs[izq], cs[der] = cs[der], cs[izq]
  QSort(cs[:izq])
  QSort(cs[(izq + 1):])
  return cs
}

/*
  Busca prioridad de candidato por BFS
*/
func (cand *Candidato) GetPrioridad(esc *Escenario){
  idIndiceV := cand.Id
  porProcesar := make([]*util.VerticeD, 0)
  porProcesar = append(porProcesar, Dirigida.Vertices[idIndiceV])
  prioridad := 0.0
  mapeo := make([]int, 0)
  aux := 1;
  for _, v := range porProcesar{
    if(ConsultaEstado(idIndiceV) == 2){
      mapeo[aux]++
    }
    aux++
    for _, u := range v.Hijos{
      porProcesar = append(porProcesar, Dirigida.Vertices[u])
    }
    porProcesar = porProcesar[1:]
  }
  aux = 1
  for _, i := range mapeo{
    if(i != 0){
      prioridad += (float64(aux) * float64(i))
    }
    aux /= 2
  }
  cand.Prioridad = prioridad
}
