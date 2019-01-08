package problema_bombero

import(
  // "fmt"
)

// Ids del conjunto que hay que salvar a toda costa
  // var PorSalvar []int
//Arboles a partir de los vertices de S-Fire
  var Arboles []Arbol

type ArbolVertice struct{
  arb int
  vert Vertice
  eval float64

}

type Candidato struct{
  Id int
  Relaciones []ArbolVertice
  Incidencias int
}

// func GetCandidatos(vecin Vecindario) []int{
//
// }
