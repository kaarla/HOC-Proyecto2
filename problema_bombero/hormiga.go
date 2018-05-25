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

func (hormiga *Hormiga) CalculaSolucion() *Solucion{
  solucion := Solucion{}
  solucion.Trayecto = hormiga.Trayecto
  solucion.Costo = hormiga.CalculaCosto()
  return &solucion
}

func (escenario *Escenario) reducePheActual(){
  escenario.PheActual = escenario.PheActual - PheReducion
}

func (hormiga *Hormiga) CalculaCosto() float64{
  quemados1 := float64(len(hormiga.Trayecto[0].Ve.GetIncendiados()))
  quemadosT := float64(len(hormiga.Trayecto[len(hormiga.Trayecto) - 1].Ve.GetIncendiados()))
  bomberosT := float64(len(hormiga.Trayecto[len(hormiga.Trayecto) - 1].Ve.GetDefendidos()))
  dano1 := quemados1 / float64(len(hormiga.Trayecto[0].Ve.Manzanas))
  danoT := quemadosT / float64(len(hormiga.Trayecto[0].Ve.Manzanas))
  d := (dano1 / danoT)
  b := (bomberosT / float64(TotalBomberos))
  return  (d * b) * float64(len(hormiga.Trayecto))
}

func (hormiga *Hormiga) AvanzaHormiga(){
  d1 := 0
  d1 = randInt(0, 2)
  nuevoEscenario := Escenario{}
  candidatos := hormiga.Actual.Ve.GetCandidatos()
  // fmt.Println("candidatosA", candidatos)
  if(len(candidatos) < 1){
    hormiga.Camina = false
    sol := hormiga.CalculaSolucion()
    sol.Trayecto[len(sol.Trayecto) - 1].Ve.PrintManzana()
    fmt.Println("Costo:", sol.Costo)
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
    //agregar arreglo para que ahi se haga el append
    hormiga.Actual.Vecinos = append(hormiga.Actual.Vecinos, nuevoEscenario)
    if(hormiga.Actual.MejorVecino != nil && hormiga.Actual.MejorVecino.Eval < nuevoEscenario.Eval){
      *hormiga.Actual.MejorVecino = nuevoEscenario
    }
    hormiga.Actual = nuevoEscenario
    nuevoEscenario.PheActual = nuevoEscenario.PheActual + Phe
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

func fin() bool{
  termino := false
  for _, h := range HormigasCaminantes{
    termino = termino || h.Camina
    fmt.Println(termino, h.Camina)
  }
  return termino
}

func randInt(min int, max int) int {
  return min + rand.Intn(max-min)
}

func CorreHeuristica(grafica string, fuegoInicial []int){
  rand.Seed(Semilla)
  vecindarioCero := VecindarioCero(grafica)
  for _, i := range fuegoInicial{
    vecindarioCero.InitFuegoEspecifico(i)
  }
  fmt.Println("-------- INICIAL ---------")
  vecindarioCero.PrintManzana()
  fmt.Println("---------------------------")
  escenarioCero := InitEscenario(vecindarioCero)
  for i := 0; i < HormigasXt; i++{
    HormigasCaminantes = append(HormigasCaminantes, *InitHormiga(i, escenarioCero))
  }
  for c:= 1; c <= 10; c++{
    for _, b := range HormigasCaminantes{
      // fmt.Println("c", c)
      if(b.Camina == true){
        // fmt.Println("b.camina", b.Camina)
        b.AvanzaHormiga()
        b.Actual.Ve.PropagaFuego()
        fmt.Println("quemados", len(b.Actual.Ve.GetIncendiados()))
        fmt.Println("no-quemados", len(b.Actual.Ve.GetASalvo()) + len(b.Actual.Ve.GetDefendidos()))
      }
    }
  }
}
