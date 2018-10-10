package problema_bombero

import (
  "fmt"
  "math/rand"
)

var Time int
var Phe float64
var PheReducion float64
var TotalBomberos int
var BomberosXt int
var HormigasXt int
var HormigasCaminantes []Hormiga
var Semilla int64
var q1 int

type Solucion struct{
  Trayecto []Escenario
  Costo float64
}

type Hormiga struct{
  Id int
  Actual Escenario
  Trayecto []Escenario
  Camina bool
}

type Escenario struct{
  Ve Vecindario
  PheActual float64
  Eval float64
  Vecinos []Escenario
  MejorVecino *Escenario
}

func InitEscenario(vecindario Vecindario) *Escenario{
  escenario := Escenario{}
  escenario.Ve= vecindario.Copia()
  escenario.PheActual = 0.0
  escenario.Eval = escenario.Ve.Evalua(TotalBomberos)
  escenario.Vecinos = nil
  escenario.MejorVecino = nil
  return &escenario
}

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
  return &solucion
}

func (escenario *Escenario) reducePheActual(){
  escenario.PheActual = escenario.PheActual - PheReducion
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
  d1 = randInt(0, 2)
  nuevoEscenario := Escenario{}
  candidatos := hormiga.Actual.Ve.GetCandidatos()
  if(len(candidatos) < 1){
    hormiga.Camina = false

    sol := hormiga.CalculaSolucion(c)
    fmt.Println("Semilla:", Semilla)
    fmt.Println("Costo:", sol.Costo)
    fmt.Println("Salvados: ", len(hormiga.Actual.Ve.GetASalvo()) + len(hormiga.Actual.Ve.GetDefendidos()))
    fmt.Println("Bomberos usados: ", len(hormiga.Actual.Ve.GetDefendidos()))
   // sol.Trayecto[len(sol.Trayecto) - 1].Ve.PrintSVG()

    return false
  }else{
    if(hormiga.Actual.Vecinos != nil && d1 == 0){
      nuevoEscenario = *hormiga.Actual.MejorVecino
    }
    d1 = randInt(0, 2)
    if(hormiga.Actual.Vecinos != nil && d1 == 1){
      nuevoEscenario = hormiga.Actual.Vecinos[(randInt(0, len(hormiga.Actual.Vecinos)))]
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

func (escenario *Escenario) copia() Escenario{
    escenarioN := Escenario{}
    escenarioN.Ve = escenario.Ve
    escenarioN.PheActual = escenario.PheActual
    escenarioN.Eval = escenario.Eval
    escenarioN.Vecinos = escenario.Vecinos
    escenarioN.MejorVecino = escenario.MejorVecino
    return escenarioN
}

func (hormiga* Hormiga) copia() Hormiga{
  hormigaN := Hormiga{}
  hormigaN.Id = hormiga.Id
  hormigaN.Actual = hormiga.Actual
  hormigaN.Trayecto = hormiga.Trayecto
  hormigaN.Camina = hormiga.Camina
  return hormigaN
}

func (hormiga *Hormiga) newEscenario(candidatos []int) Escenario{
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
      r1 = randInt(0, len(candidatos))
      if(Contiene(bomberosN, r1)){
        i--
      }else{
        bomberosN = append(bomberosN, r1)
      }
    }
  }
  for i := 0; i < len(bomberosN); i++{
    escenario.Ve.Manzanas[candidatos[bomberosN[i]]].Estado = 1
  }
  return escenario
}


func randInt(min int, max int) int {
  return min + rand.Intn(max-min)
}

func CorreHeuristica(grafica string, fuegoInicial []int){
  rand.Seed(Semilla)
  q1 = len(fuegoInicial)
  vecindarioCero := VecindarioCero(grafica)
  for _, i := range fuegoInicial{
    vecindarioCero.InitFuegoEspecifico(i)
  }
   fmt.Println("-------- INICIAL ---------")
   vecindarioCero.PrintSVG()
   fmt.Println("---------------------------")
  escenarioCero := InitEscenario(vecindarioCero)
  for i := 0; i < HormigasXt; i++{
    HormigasCaminantes = append(HormigasCaminantes, *InitHormiga(i, escenarioCero))
  }
  fin := true
  c := 1
  for fin{
    for _, b := range HormigasCaminantes{
      t := true
      if(t){
        t = b.AvanzaHormiga(c)
        b.Actual.Ve.PropagaFuego()
        fin = b.Camina
        c++
      }
    }
  }
}
