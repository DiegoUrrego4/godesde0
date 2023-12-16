package ejercicios

import (
	"strconv"
)

func ConvertStringToInt(stringValue string) (int, string) {
	// Conversión de string a Int
	intNumber, err := strconv.Atoi(stringValue)

	// Manejo del error
	if err != nil {
		return 0, "Hubo un error " + err.Error()
	}

	if intNumber > 100 {
		return intNumber, "Es mayor a 100"
	} else {
		return intNumber, "Es menor a 100"
	}
}
