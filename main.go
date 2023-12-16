package main

import (
	"fmt"

	"github.com/diegourrego4/godesde0/variables"
)

func main() {
	state, text := variables.ConvertToText(1588)
	fmt.Println("state =", state)
	fmt.Println("text =", text)
}
