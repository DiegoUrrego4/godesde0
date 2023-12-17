package arreglos_slices

import "fmt"

var tablaSlice []int = []int{2, 5, 4}
var arreglo [10]int = [10]int{6, 98, 4, 56, 734, 24, 14, 2, 4, 5}

func MuestroSlice() {
	fmt.Println(tablaSlice)
	// ! Nota: En los slices, la última posición es excluyente
	porcion := arreglo[3:]   // Slice creado desde un vector, desde la posición 3
	porcion2 := arreglo[:5]  // Slice desde el comienzo hasta la posición 5
	porcion3 := arreglo[6:7] // Slice desde la posición 6 a la 7

	fmt.Println("porción =", porcion)
	fmt.Println("porción2 =", porcion2)
	fmt.Println("porcion3 =", porcion3)

}

// Los arreglos no pueden agrandar su capacidad, los slices sí.
func Capacidad() {
	elementos := make([]int, 5, 20) // 5 elementos con 20 campos de capacidad
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))

	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}

	fmt.Printf("\nLargo %d, Capacidad %d", len(nums), cap(nums))

}
