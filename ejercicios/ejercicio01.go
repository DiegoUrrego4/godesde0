package ejercicios

import (
	"strconv"
)

func ConvertStringToInt(stringValue string) (int, string) {
	// Conversión de string a Int
	valueConvertedToInt, _ := strconv.Atoi(stringValue)

	if valueConvertedToInt > 100 {
		return valueConvertedToInt, "Es mayor a 100"
	} else {
		return valueConvertedToInt, "Es menor a 100"
	}
}
