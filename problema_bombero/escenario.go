package problema_bombero
import(
  "github.com/kaarla/HOC-Proyecto2/util"
  "sort"
  "math/rand"
  // "container/list"
  // "fmt"
)


//Estructura para el escenario
type Escenario struct{
  Ve Vecindario            //estado de la gráfica en la que se extiende el inciendio
  PheActual float64        //feromonas actuales en el escenario
  Eval float64             //evaluación del escenario dado el vecindario
  Vecinos []*Escenario      //escenarios conocidos a los que se puede llegar desde el actual en una unidad de tiempo
  MejorVecino *Escenario   //vecino con mejor evaluación
  DistanciaAe0 int         //distancia al hormiguero
}

/*
  Inicializa el vecindario vacío
*/
func NewEscenario(vecindario Vecindario) *Escenario{
  escenario := Escenario{}
  escenario.Ve= vecindario.Copia()
  escenario.PheActual = 0.0
  escenario.Eval = escenario.Ve.Evalua(TotalBomberos)
  escenario.Vecinos = nil
  escenario.MejorVecino = nil
  escenario.DistanciaAe0 = 0
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
  if(a > b){
    a, b = b, a
  }
  var estado int
  path := []int{}
    a1 := Trayectorias[a][b]
  for a != b && a != a1{
    a1 = a
    a = Trayectorias[a1][b]
    estado = escenario.Ve.ConsultaEstado(a)
    // aux := escenario.Ve.Manzanas[a]
    if(a != b){
      if(estado == 0){
        path = append(path, a)
      }else{
        return []int{}
      }
    }
  }
  return path
}

/*
crea los candidatos para un escenario
*/
func (esc *Escenario) GetCandidatos() []*Candidato{
  incendiados := esc.Ve.GetIncendiados()
  candidatosBrut := []int{}
  candidatos := []*Candidato{}
  actual := 0
  for _, s := range PorSalvar{
    for _, b := range incendiados{
      candidatosBrut = append(candidatosBrut, esc.GetTrayectoria(s, b)...)
    }
  }
  sort.Ints(candidatosBrut)

  for len(candidatosBrut) > 1{
    actual = candidatosBrut[0]
    incidencias := util.Cuenta(candidatosBrut, actual)
    if(!util.Contiene(PorSalvar, actual)){
      newCand := NewCandidato(actual, incidencias)
      newCand.FindMins(esc.Ve)
      newCand.GetPrioridad(esc);
      candidatos = append(candidatos, newCand)
    }
    candidatosBrut = append(candidatosBrut[:0], candidatosBrut[(incidencias):]...)
  }
  QSort(candidatos)
  // for _, c := range candidatos{
  //   fmt.Println("<p>c ", c.Id, "</p>")
  // }
  return candidatos
}

func CreaEscenario(candidatos []*Candidato, actual *Escenario, distancia int) *Escenario{
  rand.Seed(Semilla)
  escenario := actual.copia()
  bomberosN := []int{}
  // r1 := 0
  if(len(candidatos) <= BomberosXt){
    for i:= 0; i < len(candidatos); i++{
      bomberosN = append(bomberosN, candidatos[i].Id)
    }
  }else{
    for i := 0; i < BomberosXt; i++{
      bomberosN = append(bomberosN, candidatos[len(candidatos) - 1].Id)
      candidatos = candidatos[:len(candidatos) -1]
    }
  }
  for i := 0; i < len(bomberosN); i++{
    escenario.Ve.Manzanas[bomberosN[i]].SetEstado(1)
    // escenario.Ve.Manzanas[bomberosN[i]].Estado = 1
  }
  escenario.DistanciaAe0 = distancia
  return &escenario
}

func (esc *Escenario)EncuentraRegreso() *Escenario{
  // fmt.Println("regresa")
  regreso := esc
  dist := esc.DistanciaAe0
  for _, e := range esc.Vecinos{
    if(dist > e.DistanciaAe0){
      regreso = e
      dist = e.DistanciaAe0
    }
  }
  return regreso
}
