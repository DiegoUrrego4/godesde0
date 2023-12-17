package arreglos_slices

import "fmt"

// Los slices no tienen una cantidad de elementos, los vectores sí.
// Vector:
var tabla [10]int = [10]int{10, 0, 59}

// Matriz:
// []-> cantidad de arrays [] -> cantidad de posiciones
var matriz [2][3]int // 2 arrays con 3 posiciones

func MuestroArreglos() {
	// Asignación directa
	tabla[7] = 33
	tabla[2] = 54

	tabla2 := [10]string{"Pablo", "Juan"}
	fmt.Println(tabla)
	fmt.Println(tabla2)

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}

	matriz[1][2] = 15
	fmt.Println(matriz)

}
