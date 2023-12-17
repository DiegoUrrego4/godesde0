package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/diegourrego4/godesde0/ejercicios"
)

var fileName string = "./files/txt/table.txt"

func SaveTable() {
	var text string = ejercicios.ShowMultiplyingTable()
	file, err := os.Create(fileName)

	if err != nil {
		fmt.Println("Error al crear el archivo" + err.Error())
		return
	}

	fmt.Fprintln(file, text)
	file.Close()
}

func AddTableToFile() {
	var text string = ejercicios.ShowMultiplyingTable()
	if !Append(fileName, text) {
		fmt.Println("Error al concatenar contenido")
	}
}

func Append(fileName string, text string) bool {
	arch, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error durante el append " + err.Error())
		return false
	}
	_, err = arch.WriteString(text)
	if err != nil {
		fmt.Println("Error durante el WriteString " + err.Error())
		return false
	}

	arch.Close()
	return true
}

func ReadFile() {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Hubo un error al leer el archivo " + err.Error())
		return
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		registro := scanner.Text() // Esto devuelve cada una de las líneas del archivo
		fmt.Println("> " + registro)
	}
	file.Close()
}
