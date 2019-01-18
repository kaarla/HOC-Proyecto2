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

/*
  Compara dos candidatos
  si es mayor, devuelve 1
  si es igual, devuelve 0
  si es menor, devuelve -1
*/
func (cA *Candidato) compareTo(cB *Candidato) int{
  if(cA.DistMinS == 1 && cA.DistMinB == 1){
    return 1
  }
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

// func QSort(cs []*Candidato, p int, fin int){
//   if((fin - p + 1) < 2){
//     return
//   }
//   i := p + 1
//   j := fin
//   for i < j {
//     fmt.Println("<p>i, j  ", i, j, "</p>")
//     if(cs[i].compareTo(cs[p]) > 0 && cs[j].compareTo(cs[p]) < 0){
//       i++
//       j--
//       intercambia(cs, i, j)
//     }else if(cs[i].compareTo(cs[p]) <= 0 && cs[j].compareTo(cs[p]) >= 0){
//       i++
//       j++
//     }else if(cs[i].compareTo(cs[p]) <= 0){
//       i++
//     }else{
//       j--
//     }
//   }
//   if(cs[i].compareTo(cs[p]) >= 0){
//     i--
//   }
//   intercambia(cs, p, i)
//   if(j - p > 1){
//     QSort(cs, p, j - 1)
//   }
//   if(fin - 1 > 1){
//     QSort(cs, i + 1, fin)
//   }
// }

func QSort(cs []*Candidato) []*Candidato{
  if len(cs) < 2{
    return cs
  }
  izq, der := 0, len(cs) - 1
  p := int(len(cs) / 2)
  cs[p], cs[izq] = cs[der], cs[p]
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

// func intercambia(array []*Candidato, i int, j int){
//   aux1 := array[i]
//   aux2 := array[j]
//   array[i] = aux2
//   array[j] = aux1
// }
