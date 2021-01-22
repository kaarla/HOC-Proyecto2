package problema_bombero

//Estructura para una solución, guardo su trayecto que es un arreglo de escenarios
// y su costo
type Solucion struct{
  Trayecto []*Escenario
  Costo float64
  Factible bool
  Semilla int64
  CostoIteraciones float64
  CostoBomberos float64
}

func NewSolucion() *Solucion{
  sol := Solucion{}
  sol.Costo = 10000.0
  sol.Factible = false
  sol.Semilla = Semilla
  return &sol
}

func (solucion *Solucion) CalculaCosto(c int, trayectoria []*Escenario){
  quemados1 := float64(q1)
  quemadosT := float64(len(trayectoria[len(trayectoria) - 1].Ve.GetIncendiados()))
  bomberosT := float64(len(trayectoria[len(trayectoria) - 1].Ve.GetDefendidos()))
  dano1 := quemados1 / float64(c)
  danoT := quemadosT / float64(c)
  d := (danoT - dano1)
  b := (bomberosT / float64(TotalBomberos)) * (bomberosT / float64(TotalBomberos))
  solucion.Costo = (d * b) * float64(c / 10)
  solucion.CostoBomberos = d * b
  solucion.CostoIteraciones = (d * b) * float64(c / 10)
}


func CalculaSolucion(c int, trayectoria []*Escenario, actual *Escenario) *Solucion{
  solucion := NewSolucion()
  // fmt.Println("<p>len de trayectoria en calculaS", len(trayectoria), "</p>")
  solucion.Trayecto = trayectoria
  solucion.CalculaCosto(c, trayectoria)
  factible := true
  aux := false
  for _, ps := range PorSalvar{
      if(actual.Ve.ConsultaEstado(ps) == 0){
      aux = true
    }
      factible = factible && aux
  }
  solucion.Factible = factible
  return solucion
}
