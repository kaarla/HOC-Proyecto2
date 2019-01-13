package problema_bombero
import(
  // "fmt"
  "github.com/kaarla/HOC-Proyecto2/util"
)

//Modela un candidato a defender
type Candidato struct{
  Id int
  NumTrayectorias int
  DistMinB int
  DistMinS int
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
  return cand
}

/*
  Encuentra el Id de los vertices mas cercanos en fuego y por salvar del candidato
*/
func (c *Candidato) FindMins(dists []int, manzanas []Manzana){
  // fmt.Println("len de manzanas en FindMins, ", len(manzanas))
  // fmt.Println("len de distancias en FindMins, ", len(dists))

  var minS int
  var minB int
  auxS := false
  auxB := false
  i := 0
  for !(auxB && auxS){

    if(manzanas[i].Estado == 2 && !auxB){
      minB = manzanas[i].Id
      c.DistMinB = minB
      auxB = true
    }
    if(util.Contiene(PorSalvar, manzanas[i].Id) && !auxS){
      minS = manzanas[i].Id
      c.DistMinS = minS
      auxS = true
    }
    i++
  }
}
