package main

import (
	"fmt"
	"github.com/diegourrego4/godesde0/ejercicios"
	"github.com/diegourrego4/godesde0/variables"
	"runtime"
)

func main() {
	state, text := variables.ConvertToText(1588)

	fmt.Println("state =", state)
	fmt.Println("text =", text)

	os := runtime.GOOS

	fmt.Println("Sistema operativo =", os)

	if os == "Linux." {
		fmt.Println("Sistema operativo Linux!")
	} else if os == "darwin" {
		fmt.Println("Sistema operativo Mac!")
	} else {
		fmt.Println("Sistema operativo Windows!")
	}

	switch os {
	case "Linux .":
		fmt.Println("Sistema operativo Linux!")
	case "darwin":
		fmt.Println("Sistema operativo Mac!")
	default:
		fmt.Printf("%s \n", os)
	}

	numberConverted, message := ejercicios.ConvertStringToInt("80")
	fmt.Println("numberConverted =", numberConverted)
	fmt.Println("message =", message)

}
