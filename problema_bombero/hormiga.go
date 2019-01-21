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
  sol.Costo = 10000.0
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
  fmt.Println("<p>len de trayectoria en calculaS", len(hormiga.Trayecto), "</p>")
  solucion.Trayecto = hormiga.Trayecto
  solucion.Costo = hormiga.CalculaCosto(c)
  factible := true
  temp := false
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
  porQuemar := hormiga.Actual.Ve.GetPorQuemar()

  hormiga.Actual.Ve.PrintSVG()
  if(len(porQuemar) < 1){
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
    return false
  }
}

func (hormiga *Hormiga) Regresa() bool{
  fmt.Println("regresa hormiga en index, ", hormiga.Index)
  if(hormiga.Index <= 0){
    hormiga.Camina = false
    return true
  }
  hormiga.Actual = hormiga.Trayecto[hormiga.Index]
  hormiga.Actual.PheActual = hormiga.Actual.PheActual + Phe
  hormiga.Index = hormiga.Index - 1
  return false
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
    escenario.Ve.Manzanas[bomberosN[i]].Estado = 1
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
   vecindarioCero.PrintSVG()
   // fmt.Println("---------------------------")
  escenarioCero := InitEscenario(vecindarioCero)
  fin := true
  ciclos := 0
  cuentaGeneraciones := 0
  cuentaTerminadas := 0
  mejorSol := NewSolucion()
  idHormiga := 0
  for fin{
    if(cuentaGeneraciones < generaciones){
      for i := 0; i < HormigasXt; i++{
        // fmt.Println("generando hormiga")
        HormigasCaminantes = append(HormigasCaminantes, *InitHormiga(i + (ciclos * HormigasXt), escenarioCero))
      }
      cuentaGeneraciones++
    }
    termino := false
    for i, b := range HormigasCaminantes{
      // fmt.Println("hormigas avanzan, ", i)
      if !termino{
        termino = b.AvanzaHormiga(ciclos)
        // fmt.Println("no termino ", i, termino)
        if(termino){
          cuentaTerminadas++
          // fmt.Println("<p>terminaron, ", cuentaTerminadas, "len trayecto", len(b.Trayecto), "</p>")
          solActual := b.CalculaSolucion(ciclos)
          if(mejorSol.Costo >= solActual.Costo && solActual.Factible){
            mejorSol = solActual
            idHormiga = b.Id
          }
        }
        if(cuentaTerminadas == (generaciones * HormigasXt)){
          fin = false
        }
        b.Actual.Ve.PropagaFuego()
        HormigasCaminantes[i] = b
      }
    }
    ciclos++
  }
  fmt.Println("<p>Seed:", Semilla, "</p>")
  // fmt.Println("<p>Saved: ", len(mejorSol.Trayecto[len(mejorSol.Trayecto) - 1].Ve.GetASalvo()) + len(mejorSol.Trayecto[len(mejorSol.Trayecto) - 1].Ve.GetDefendidos()), "</p>")
  // fmt.Println("<p>Total of firefighters: ", len(mejorSol.Trayecto[len(mejorSol.Trayecto) - 1].Ve.GetDefendidos()), "</p>")
  fmt.Println("<p>Firefighters in each t: ", BomberosXt, "</p>")
  fmt.Println("<p>Cost:", mejorSol.Costo, "</p>")
  fmt.Println("<p>Fact:", mejorSol.Factible, "</p>")
  fmt.Println("<p>Pasos:", len(mejorSol.Trayecto), "</p>")
  fmt.Println("<p>HormigaId:", idHormiga, "</p>")
  fmt.Println("<p>Pasos HORMIGA:", len(HormigasCaminantes[idHormiga].Trayecto), "</p>")
  mejorSol.Trayecto = HormigasCaminantes[idHormiga].Trayecto
  for i, es := range mejorSol.Trayecto{
    fmt.Println("<p>-----------------------------------</p>")
    fmt.Println("Tiempo", i + 1)
    es.Ve.PrintSVG()
  }
}
