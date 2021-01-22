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
var HormigasExploradoras []Hormiga
//semilla que se usará para inicializar el random
var Semilla int64
//número de vértices que se incendiarán en t = 1
var q1 int
//Ids del conjunto que hay que salvar a toda costa
var PorSalvar []int
//Grafica dirigida sobre la que va a trabajar
var Dirigida *util.Dirigida

var NumVertices int


func CorreHeuristica(fuegoInicial []int){
  Dirigida = util.CreaDirigida(PorSalvar, NumVertices)
  rand.Seed(Semilla)
  generaciones := 3
  q1 = len(fuegoInicial)
  vecindarioCero := VecindarioCero()
  fmt.Println("len de V cero", len(vecindarioCero.Manzanas))
  for _, i := range fuegoInicial{
    vecindarioCero.InitFuegoEspecifico(i)
  }

  escenarioCero := NewEscenario(vecindarioCero)
  fmt.Println("len al crear escenario, ", len(escenarioCero.Ve.Manzanas))
  fin := true
  ciclos := 0
  cuentaGeneraciones := 0
  cuentaTerminadas := 0
  mejorSol := NewSolucion()
  idHormiga := 0

  for fin{
    if(cuentaGeneraciones < generaciones){
      for i := 0; i < HormigasXt; i++{
        HormigasExploradoras = append(HormigasExploradoras, *newHormiga(i + (ciclos * HormigasXt), escenarioCero))
      }
      cuentaGeneraciones++
    }
    termino := false
    for i, b := range HormigasExploradoras{
      if !termino{
        termino = b.avanza(ciclos)
        if(termino){
          cuentaTerminadas++
          solActual := CalculaSolucion(ciclos, b.Trayecto, b.Actual)
          if(mejorSol.Costo >= solActual.Costo /*&& solActual.Factible*/){
            mejorSol = solActual
            idHormiga = b.Id
          }
        }
        if(cuentaTerminadas == (generaciones * HormigasXt)){
          fin = false
        }
        b.Actual.Ve.PropagaFuego()
        HormigasExploradoras[i] = b
      }
    }
    ciclos++
  }
  // for _, h := range HormigasExploradoras{
  //   fmt.Println("Hormiga", h.Id, "trayectoria de: ", len(h.Trayecto))
  // }
  printResultado(mejorSol, idHormiga)
}

func printResultado(mejorSol *Solucion, idHormiga int){
  fmt.Println("<p>Seed:", Semilla, "</p>")
  // fmt.Println("<p>Saved: ", len(mejorSol.Trayecto[len(mejorSol.Trayecto) - 1].Ve.GetASalvo()) + len(mejorSol.Trayecto[len(mejorSol.Trayecto) - 1].Ve.GetDefendidos()), "</p>")
  // fmt.Println("<p>Total of firefighters: ", len(mejorSol.Trayecto[len(mejorSol.Trayecto) - 1].Ve.GetDefendidos()), "</p>")
  // fmt.Println("<p>Firefighters in each t: ", BomberosXt, "</p>")
  fmt.Println("<p>Cost:", mejorSol.Costo, "</p>")
  fmt.Println("<p>Fact:", mejorSol.Factible, "</p>")
  // fmt.Println("<p>Pasos:", len(mejorSol.Trayecto), "</p>")
  // fmt.Println("<p>HormigaId:", idHormiga, "</p>")
  // fmt.Println("<p>Pasos HORMIGA:", len(HormigasExploradoras[idHormiga].Trayecto), "</p>")
  mejorSol.Trayecto = HormigasExploradoras[idHormiga].Trayecto
  // for i, es := range mejorSol.Trayecto{
  //   fmt.Println("<p>-----------------------------------</p>")
  //   fmt.Println("Tiempo", i + 1)
  //   es.Ve.PrintSVG()
  // }
}
