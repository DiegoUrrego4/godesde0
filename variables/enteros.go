package variables

import "fmt"

// Esta función debe iniciar en mayúscula, pues es necesario para que se importe
func ShowIntegers() {
	// Declaración de variables. Existen dos tipos:
	// Por asignación
	var intComun int
	intde32 := int32(10)
	intde64 := int64(100)

	fmt.Println("intcomun = ", intComun)
	fmt.Println("intde32 = ", intde32)
	fmt.Println("intde64 = ", intde64)
}
