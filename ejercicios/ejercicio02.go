package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ShowMultiplyingTable() {
	var number int
	var err error

	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Ingresa el número para multiplicar:")
		if scanner.Scan() {
			number, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Error. Ingresa un número válido")
				continue
			} else {
				break
			}
		}
	}

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d \n", number, i, number*i)
	}

}
