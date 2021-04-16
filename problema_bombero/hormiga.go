package problema_bombero

import(
  // "fmt"
  "github.com/kaarla/HOC-Proyecto2/util"
)

//Estructura para la hormiga
type Hormiga struct{
  Id int                  //id para identificarla
  Actual *Escenario        //escenario en el que se encuentra
  Trayecto []*Escenario    //trayectoria que siguió hasta el momento
  Camina bool             //booleano para saber si ya llegó a la condición de paro
  Ida bool                // true si va, false si regresa
  Index int               // indice del escenario de la trayectoria en el que va
  Pasos int
}

/*
  Inicializa una hormiga con un escenario.
*/
func newHormiga(id int, escenario *Escenario) *Hormiga{
  hormiga := Hormiga{}
  hormiga.Id = id
  hormiga.Actual = escenario
  hormiga.Trayecto = append(hormiga.Trayecto, escenario)
  hormiga.Camina = true
  return &hormiga
}

func (hormiga *Hormiga) avanza(ciclo int) bool{
  d1 := 0
  d1 = util.RandInt(0, 2)
  nuevoEscenario := &Escenario{}
  candidatos := hormiga.Actual.GetCandidatos()
  porQuemar := hormiga.Actual.Ve.GetPorQuemar()
  // fmt.Println("INDICE: ", hormiga.Index)

  // hormiga.Actual.Ve.PrintSVG()
  if(len(porQuemar) < 1){
    if (hormiga.Ida){
      hormiga.Ida = false
    }
    return hormiga.regresa()
  }else{
    if(hormiga.Actual.Vecinos != nil && d1 == 0){
      nuevoEscenario = hormiga.Actual.MejorVecino
    }
    d1 = util.RandInt(0, 2)
    if(hormiga.Actual.Vecinos != nil && d1 == 1){
      nuevoEscenario = hormiga.Actual.Vecinos[(util.RandInt(0, len(hormiga.Actual.Vecinos)))]
    }else{
      nuevoEscenario = CreaEscenario(candidatos, hormiga.Actual, len(hormiga.Trayecto))
    }
    hormiga.Actual.Vecinos = append(hormiga.Actual.Vecinos, nuevoEscenario)
    hormiga.Trayecto = append(hormiga.Trayecto, nuevoEscenario)

    if(hormiga.Actual.MejorVecino != nil && hormiga.Actual.MejorVecino.Eval < nuevoEscenario.Eval){
      hormiga.Actual.MejorVecino = nuevoEscenario
    }
    hormiga.Actual = nuevoEscenario
    nuevoEscenario.PheActual = nuevoEscenario.PheActual + Phe
    hormiga.Index += 1
    hormiga.Pasos += 1
    return false
  }
}

func (hormiga *Hormiga) regresa() bool{
  // fmt.Println("regresa hormiga,", hormiga.Id,  "en index, ", hormiga.Index)
  if(hormiga.Index <= 0){
    hormiga.Camina = false
    return true
  }
  hormiga.Actual = hormiga.Actual.EncuentraRegreso()
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
