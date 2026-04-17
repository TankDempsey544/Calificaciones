package main

import "fmt"


const APROBADO float64 = 6.0
const EXCELENTE float64 = 9.0
const ASISTENCIA_MINIMA int = 75

func main() {

	var nombre0 string = "Ana García"
	var nombre1 string = "Luis Pérez"
	var nombre2 string = "María López"
	var nombre3 string = "David Romero"
	var nombre4 string = "Sofia Martínez"
	var nombre5 string = "Carlos Jiménez"
	var nombre6 string = "Isabella Torres"
	var nombre7 string = "Miguel Flores"
	var nombre8 string = "Valentina Ruiz"
	var nombre9 string = "Andrés Castro"

	asistencia0 := 95
	asistencia1 := 80
	asistencia2 := 60
	asistencia3 := 72
	asistencia4 := 88
	asistencia5 := 55
	asistencia6 := 91
	asistencia7 := 70
	asistencia8 := 85
	asistencia9 := 40

	var calificaciones0 [5]float64 = [5]float64{9.5, 9.2, 8.8, 9.7, 9.1}
	var calificaciones1 [5]float64 = [5]float64{7.0, 6.5, 7.8, 6.9, 7.2}
	var calificaciones2 [5]float64 = [5]float64{8.0, 7.5, 8.2, 7.9, 8.1}
	var calificaciones3 [5]float64 = [5]float64{4.2, 3.8, 5.1, 4.7, 4.5}
	var calificaciones4 [5]float64 = [5]float64{9.1, 8.7, 9.3, 8.9, 9.0}
	var calificaciones5 [5]float64 = [5]float64{5.5, 6.0, 5.8, 6.2, 5.9}
	var calificaciones6 [5]float64 = [5]float64{7.3, 8.1, 7.6, 8.0, 7.8}
	var calificaciones7 [5]float64 = [5]float64{6.1, 5.9, 6.4, 6.0, 6.3}
	var calificaciones8 [5]float64 = [5]float64{8.5, 9.0, 8.3, 8.8, 8.6}
	var calificaciones9 [5]float64 = [5]float64{3.5, 4.0, 3.8, 4.2, 3.9}

	var nombres [10]string
	nombres[0] = nombre0
	nombres[1] = nombre1
	nombres[2] = nombre2
	nombres[3] = nombre3
	nombres[4] = nombre4
	nombres[5] = nombre5
	nombres[6] = nombre6
	nombres[7] = nombre7
	nombres[8] = nombre8
	nombres[9] = nombre9

	var asistencias [10]int
	asistencias[0] = asistencia0
	asistencias[1] = asistencia1
	asistencias[2] = asistencia2
	asistencias[3] = asistencia3
	asistencias[4] = asistencia4
	asistencias[5] = asistencia5
	asistencias[6] = asistencia6
	asistencias[7] = asistencia7
	asistencias[8] = asistencia8
	asistencias[9] = asistencia9

	var todasLasCalificaciones [10][5]float64
	todasLasCalificaciones[0] = calificaciones0
	todasLasCalificaciones[1] = calificaciones1
	todasLasCalificaciones[2] = calificaciones2
	todasLasCalificaciones[3] = calificaciones3
	todasLasCalificaciones[4] = calificaciones4
	todasLasCalificaciones[5] = calificaciones5
	todasLasCalificaciones[6] = calificaciones6
	todasLasCalificaciones[7] = calificaciones7
	todasLasCalificaciones[8] = calificaciones8
	todasLasCalificaciones[9] = calificaciones9

	var sumaGeneral float64 = 0
	var totalAprobados int = 0
	var mejorNota float64 = 0.0
	var peorNota float64 = 10.0
	var nombreMejor string = ""
	var nombrePeor string = ""

	fmt.Println("=== SISTEMA DE CALIFICACIONES ===")
	fmt.Println()

	for i := 0; i < 10; i++ {

		var suma float64 = 0
		for indice, nota := range todasLasCalificaciones[i] {
			fmt.Println("  calificacion", indice+1, "=", nota)
			suma = suma + nota
		}

		promedio := suma / 5

		if promedio > mejorNota {
			mejorNota = promedio
			nombreMejor = nombres[i]
		}
		if promedio < peorNota {
			peorNota = promedio
			nombrePeor = nombres[i]
		}

		if promedio >= EXCELENTE {

			fmt.Println(nombres[i] + ": Excelente")
			totalAprobados = totalAprobados + 1

		} else if promedio >= APROBADO {

			if asistencias[i] >= ASISTENCIA_MINIMA {
				fmt.Println(nombres[i] + ": Aprobado")
				totalAprobados = totalAprobados + 1
			} else {
				fmt.Println(nombres[i] + ": Aprobado con baja asistencia")
			}

		} else {

			if asistencias[i] < ASISTENCIA_MINIMA {
				fmt.Println(nombres[i] + ": Reprobado (nota y asistencia insuficientes)")
			} else {
				fmt.Println(nombres[i] + ": Reprobado")
			}

		}
		var letra string
		switch {
		case promedio >= 9.0:
			letra = "A+"
		case promedio >= 8.0:
			letra = "A"
		case promedio >= 7.0:
			letra = "B"
		case promedio >= 6.0:
			letra = "C"
		case promedio >= 5.0, promedio >= 4.0:
			letra = "D"
		default:
			letra = "F"
		}

		fmt.Println("  -> Promedio:", fmt.Sprintf("%.2f", promedio), "| Letra:", letra, "| Asistencia:", asistencias[i], "%")
		fmt.Println()

		sumaGeneral = sumaGeneral + promedio
	}

	var lineaSeparadora string = ""
	var contador int = 0
	for contador < 33 {
		lineaSeparadora = lineaSeparadora + "="
		contador = contador + 1
	}

	var promedioGeneral float64 = sumaGeneral / 10
	var porcentajeAprobados float64 = (float64(totalAprobados) / 10) * 100

	fmt.Println(lineaSeparadora)
	fmt.Println("=== ESTADÍSTICAS ===")
	fmt.Println(lineaSeparadora)
	fmt.Printf("Promedio general   : %.1f\n", promedioGeneral)
	fmt.Printf("Mejor calificación : %.2f (%s)\n", mejorNota, nombreMejor)
	fmt.Printf("Peor calificación  : %.2f (%s)\n", peorNota, nombrePeor)
	fmt.Printf("Aprobados          : %.0f%%\n", porcentajeAprobados)
}
