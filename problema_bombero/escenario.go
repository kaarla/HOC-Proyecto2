package problema_bombero
import(
  "github.com/kaarla/HOC-Proyecto2/util"
  "sort"
  // "fmt"
)

//Estructura para el escenario
type Escenario struct{
  Ve Vecindario            //estado de la gráfica en la que se extiende el inciendio
  PheActual float64        //feromonas actuales en el escenario
  Eval float64             //evaluación del escenario dado el vecindario
  Vecinos []Escenario      //escenarios conocidos a los que se puede llegar desde el actual en una unidad de tiempo
  MejorVecino *Escenario   //vecino con mejor evaluación
}

/*
  Inicializa el vecindario vacío
*/
func InitEscenario(vecindario Vecindario) *Escenario{
  escenario := Escenario{}
  escenario.Ve= vecindario.Copia()
  escenario.PheActual = 0.0
  escenario.Eval = escenario.Ve.Evalua(TotalBomberos)
  escenario.Vecinos = nil
  escenario.MejorVecino = nil
  return &escenario
}

func (escenario *Escenario) reducePheActual(){
  escenario.PheActual = escenario.PheActual - PheReducion
}

func (escenario *Escenario) copia() Escenario{
  escenarioN := Escenario{}
  escenarioN.Ve = escenario.Ve
  escenarioN.PheActual = escenario.PheActual
  escenarioN.Eval = escenario.Eval
  escenarioN.Vecinos = escenario.Vecinos
  escenarioN.MejorVecino = escenario.MejorVecino
  return escenarioN
}

/*
  Obtiene la trayectoria entre 2 vertices
*/
func (escenario *Escenario) GetTrayectoria(a int, b int) []int{
  path := []int{}
    a1 := Trayectorias[a][b]
    path = append(path, a1)
  for a != b && a != a1{
    a1 = a
    a = Trayectorias[a1][b]
    aux := escenario.Ve.Manzanas[a]
    if(a != b){
      if(aux.Estado == 0){
        path = append(path, a)
      }
    }
  }
  return path
}

/*
  crea los candidatos para un escenario
*/
func (esc *Escenario) GetCandidatos() []*Candidato{
  // fmt.Println("entra GetCandidatos")
  incendiados := esc.Ve.GetIncendiados()
  candidatosBrut := []int{}
  candidatos := []*Candidato{}
  actual := 0
  for _, s := range PorSalvar{
    // fmt.Println("<p>*POR SALVAR:", s, "</p>")
    for _, b := range incendiados{
      // fmt.Println("<p>INCENDIADO:", b, "</p>")
      candidatosBrut = append(candidatosBrut, esc.GetTrayectoria(s, b)...)
    }
  }
  sort.Ints(candidatosBrut)

  // fmt.Println("<p>candBrut", candidatosBrut, "</p>")
  for len(candidatosBrut) > 1{
    actual = candidatosBrut[0]
    incidencias := util.Cuenta(candidatosBrut, actual)
    newCand := NewCandidato(actual, incidencias)
    newCand.FindMins(esc.Ve.Mapa[actual], esc.Ve.Manzanas)
    candidatos = append(candidatos, newCand)
    candidatosBrut = append(candidatosBrut[:0], candidatosBrut[(incidencias):]...)
  }
  if(len(candidatosBrut) == 1){
    newCand := NewCandidato(candidatosBrut[0], 1)
    candidatos = append(candidatos, newCand)
  }
  // for _, c := range candidatos{
    // fmt.Println("<p>id", c.Id, "</p>")
    // fmt.Println("<p>tray", c.NumTrayectorias, "</p>")
  // }
  // fmt.Println("<p>id---------</p>")

  QSort(candidatos)
  // for _, c := range candidatos{
    // fmt.Println("<p>id", c.Id, "</p>")
    // fmt.Println("<p>tray", c.NumTrayectorias, "</p>")
  // }
  return candidatos
}
