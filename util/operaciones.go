package util

import(
  "sort"
  "math/rand"
)

/*
  promedio de distancias
*/
func getDistPromedio(mapa [][]int) float64{
  cont := 0.0
  suma := 0.0
  for i := 0; i < len(mapa); i++{
    for j := 0; j < len(mapa); j++{
      suma += float64(mapa[i][j])
      cont += 1.0
    }
  }
  return suma / cont
}

/*
  promedio de distancias de un vertice
*/
func getProm(i int, mapa [][]int) float64{
  cont := 0.0
  suma := 0.0
  for j := 0; j < len(mapa); j++{
    suma += float64(mapa[i][j])
    cont += 1.0
  }
  return suma / cont
}

/*
  moda de las distancias de un vertice
*/
func GetModa(i int, mapa [][]int) (int, int, int){
  aux := OrdenaDistancias(i, mapa)

  countAnterior, count := 0, 0
  actual, moda := aux[0], aux[0]
  for j := 0; j < len(aux); j++{
    if(aux[j] == actual){
      count++
      if(count >= countAnterior){
        moda = actual
        countAnterior = count
      }
    }else{
      actual = aux[j]
      count = 1
    }
  }
  return i, moda, countAnterior
}

/*
  Ordena las distancias de un vertice
*/
func OrdenaDistancias(i int, mapa [][]int) []int{
  aux := make([]int, len(mapa))
  for j := 0; j < len(mapa); j++{
    aux[j] = mapa[i][j]
  }
  sort.Ints(aux)
  return aux
}

/*
  encuentra la distancia mayor de un vertice
*/
func FindMayor(dists []int) int{
  return dists[(len(dists) - 1)]
}

/*
  Revisa si un elemento está en un arreglo.
*/
func Contiene(a []int, e int) bool{
  for _, b := range a{
    if b == e{
      return true
    }
  }
  return false
}

func RandInt(min int, max int) int {
  return min + rand.Intn(max-min)
}
