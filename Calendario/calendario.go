package main

import (
	"fmt"
	"time"
)

func generarCalendario(inicio, fin time.Time) []string {
	var fechas []string
	for d := inicio; !d.After(fin); d = d.AddDate(0, 0, 1) {
		fechas = append(fechas, d.Format("Monday, 02 January 2006"))
	}
	return fechas
}

func main() {
	fechaInicio := time.Date(2025, 6, 1, 0, 0, 0, 0, time.UTC)
	fechaFin := time.Date(2025, 8, 31, 0, 0, 0, 0, time.UTC)

	fechas := generarCalendario(fechaInicio, fechaFin)

	fmt.Println("Calendario de junio a agosto de 2025:")
	for _, fecha := range fechas {
		fmt.Println(fecha)
	}
}
