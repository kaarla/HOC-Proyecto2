package problema_bombero

import(
  "fmt"
  "math/rand"
  "github.com/kaarla/HOC-Proyecto2/util"
)

var Time int
//rastro que dejará la hormiga al pasar
var Phe float64
//cuánto se reducirá el rastro de feromonas por cada unidad de tiempo
var PheReducion float64
//cuántos bomberos se utilizaron en una solución
var TotalBomberos int
//cuántos bomberos se pueden asignar por unidad de tiempo
var BomberosXt int
//cuantas hormigas nuevas salen del origen por cada unidad de tiempo
var HormigasXt int
//arreglo para guardar a las hormigas que actualmente están en la heurística
var HormigasCaminantes []Hormiga
//semilla que se usará para inicializar el random
var Semilla int64
//número de vértices que se incendiarán en t = 1
var q1 int
//Ids del conjunto que hay que salvar a toda costa
var PorSalvar []int


//Estructura para una solución, guardo su trayecto que es un arreglo de escenarios
// y su costo
type Solucion struct{
  Trayecto []Escenario
  Costo float64
  Factible bool
}

//Estructura para la hormiga
type Hormiga struct{
  Id int                  //id para identificarla
  Actual Escenario        //escenario en el que se encuentra
  Trayecto []Escenario    //trayectoria que siguió hasta el momento
  Camina bool             //booleano para saber si ya llegó a la condición de paro
  Ida bool                // true si va, false si regresa
  Index int               // indice del escenario de la trayectoria en el que va
}

func NewSolucion() *Solucion{
  sol := Solucion{}
  sol.Costo = 1.0
  sol.Factible = false
  return &sol
}

/*
  Inicializa una hormiga con un escenario.
*/
func InitHormiga(id int, escenario *Escenario) *Hormiga{
  hormiga := Hormiga{}
  hormiga.Id = id
  hormiga.Actual = *escenario
  hormiga.Trayecto = append(hormiga.Trayecto, *escenario)
  hormiga.Camina = true
  return &hormiga
}

func (hormiga *Hormiga) CalculaSolucion(c int) *Solucion{
  solucion := Solucion{}
  solucion.Trayecto = hormiga.Trayecto
  solucion.Costo = hormiga.CalculaCosto(c)
  factible := true
  temp := true
  for _, ps := range PorSalvar{
    if(hormiga.Actual.Ve.Manzanas[ps].Estado == 0){
      temp = true
    }
      factible = factible && temp
  }
  solucion.Factible = factible
  return &solucion
}

func (hormiga *Hormiga) CalculaCosto(c int) float64{
  quemados1 := float64(q1)
  quemadosT := float64(len(hormiga.Trayecto[len(hormiga.Trayecto) - 1].Ve.GetIncendiados()))
  bomberosT := float64(len(hormiga.Trayecto[len(hormiga.Trayecto) - 1].Ve.GetDefendidos()))
  dano1 := quemados1 / float64(c)
  danoT := quemadosT / float64(c)
  d := (danoT - dano1)
  b := (bomberosT / float64(TotalBomberos)) * (bomberosT / float64(TotalBomberos))
  return  (d * b) * float64(c)
}

func (hormiga *Hormiga) AvanzaHormiga(c int) bool{
  d1 := 0
  d1 = util.RandInt(0, 2)
  nuevoEscenario := Escenario{}
  candidatos := hormiga.Actual.GetCandidatos()
  if(len(candidatos) < 1){
    if (hormiga.Ida){
      hormiga.Index = len(hormiga.Trayecto) - 2
      hormiga.Ida = false
    }

    return hormiga.Regresa()
  }else{
    if(hormiga.Actual.Vecinos != nil && d1 == 0){
      nuevoEscenario = *hormiga.Actual.MejorVecino
    }
    d1 = util.RandInt(0, 2)
    if(hormiga.Actual.Vecinos != nil && d1 == 1){
      nuevoEscenario = hormiga.Actual.Vecinos[(util.RandInt(0, len(hormiga.Actual.Vecinos)))]
    }else{
      nuevoEscenario = hormiga.newEscenario(candidatos)
    }
    hormiga.Actual.Vecinos = append(hormiga.Actual.Vecinos, nuevoEscenario)
    hormiga.Trayecto = append(hormiga.Trayecto, nuevoEscenario)

    if(hormiga.Actual.MejorVecino != nil && hormiga.Actual.MejorVecino.Eval < nuevoEscenario.Eval){
      *hormiga.Actual.MejorVecino = nuevoEscenario
    }
    hormiga.Actual = nuevoEscenario
    nuevoEscenario.PheActual = nuevoEscenario.PheActual + Phe
    return true
  }
}

func (hormiga *Hormiga) Regresa() bool{
  if(hormiga.Index <= 0){
    hormiga.Camina = false
    return false
  }
  hormiga.Actual = hormiga.Trayecto[hormiga.Index]
  hormiga.Actual.PheActual = hormiga.Actual.PheActual + Phe
  hormiga.Index = hormiga.Index - 1
  return true
}

func (hormiga* Hormiga) copia() Hormiga{
  hormigaN := Hormiga{}
  hormigaN.Id = hormiga.Id
  hormigaN.Actual = hormiga.Actual
  hormigaN.Trayecto = hormiga.Trayecto
  hormigaN.Camina = hormiga.Camina
  return hormigaN
}

func (hormiga *Hormiga) newEscenario(candidatos []*Candidato) Escenario{
  rand.Seed(Semilla)
  escenario := hormiga.Actual.copia()
  bomberosN := []int{}
  r1 := 0
  if(len(candidatos) <= BomberosXt){
    for i:= 0; i < len(candidatos); i++{
      bomberosN = append(bomberosN, i)
    }
  }else{
    for i:= 0; i < BomberosXt; i++{
      r1 = util.RandInt(0, len(candidatos))
      if(util.Contiene(bomberosN, r1)){
        i--
      }else{
        bomberosN = append(bomberosN, r1)
      }
    }
  }
  for i := 0; i < len(bomberosN); i++{
    escenario.Ve.Manzanas[candidatos[bomberosN[i]].Id].Estado = 1
  }
  return escenario
}

func CorreHeuristica(grafica string, fuegoInicial []int){
  rand.Seed(Semilla)
  generaciones := 3
  q1 = len(fuegoInicial)
  vecindarioCero := VecindarioCero(grafica)
  for _, i := range fuegoInicial{
    vecindarioCero.InitFuegoEspecifico(i)
  }
   // fmt.Println("-------- INICIAL ---------")
   // fmt.Println("---------------------------")
  escenarioCero := InitEscenario(vecindarioCero)
  fin := true
  ciclos := 0
  cuentaGeneraciones := 0
  cuentaTerminadas := 0
  mejorSol := NewSolucion()
  for fin{
    if(cuentaGeneraciones < generaciones){
      for i := 0; i < HormigasXt; i++{
        HormigasCaminantes = append(HormigasCaminantes, *InitHormiga(i + (ciclos * HormigasXt), escenarioCero))
      }
      cuentaGeneraciones++
    }
    t := true
    for i, b := range HormigasCaminantes{
      if t{
        t = b.AvanzaHormiga(ciclos)
        if(!t){
          cuentaTerminadas++
          solActual := b.CalculaSolucion(ciclos)
          if(mejorSol.Costo > solActual.Costo && solActual.Factible){
          // if(b.Id == 0){
            mejorSol = solActual
          }
        }
        if(cuentaTerminadas == generaciones * HormigasXt){
          fin = false
        }
        b.Actual.Ve.PropagaFuego()
        HormigasCaminantes[i] = b
      }
    }
    ciclos++
  }
  fmt.Println("<p>Seed:", Semilla, "</p>")
  fmt.Println("<p>Saved: ", len(mejorSol.Trayecto[len(mejorSol.Trayecto) - 1].Ve.GetASalvo()) + len(mejorSol.Trayecto[len(mejorSol.Trayecto) - 1].Ve.GetDefendidos()), "</p>")
  fmt.Println("<p>Total of firefighters: ", len(mejorSol.Trayecto[len(mejorSol.Trayecto) - 1].Ve.GetDefendidos()), "</p>")
  fmt.Println("<p>Firefighters in each t: ", BomberosXt, "</p>")
  fmt.Println("<p>Cost:", mejorSol.Costo, "</p>")
  fmt.Println("<p>Fact:", mejorSol.Factible, "</p>")
  for i, ve := range mejorSol.Trayecto{
    fmt.Println("i", i)
    ve.Ve.PrintSVG()
  }
  // fmt.Println("long", len(mejorSol.Trayecto))
}
