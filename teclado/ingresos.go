package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var number1 int
var number2 int
var legend string
var err error

func InsertNumbers() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese número 1: ")
	if scanner.Scan() {
		number1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			// Para abortar la aplicación se puede usar un panic
			panic("El dato ingresado es incorrecto " + err.Error())
		}
	}

	fmt.Println("Ingrese número 2: ")
	if scanner.Scan() {
		number2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			// Para abortar la aplicación se puede usar un panic
			panic("El dato ingresado es incorrecto " + err.Error())
		}
	}

	fmt.Println("Ingrese leyenda: ")
	if scanner.Scan() {
		legend = scanner.Text()
	}
	fmt.Println(legend, number1*number2)
}
