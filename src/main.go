package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	value, err := strconv.ParseFloat(os.Args[1], 64)

	if err != nil {
		fmt.Println("Valor informado inválido")
		return
	} else if len(os.Args) != 4 {
		fmt.Println("ERRO: falta de argumentos")
		return
	}

	ConvertTemp(value)
}

func ConvertTemp(value float64) {
	args1, args2 := os.Args[2], os.Args[3]
	if args1 == "-C" && args2 == "-F" {
		value = (value * 9 / 5) + 32
		fmt.Println(value)
	} else if args1 == "-F" && args2 == "-C" {
		value = (value - 32) * 5 / 9
		fmt.Println(value)
	} else if args1 == "-C" && args2 == "-K" {
		value = value + 273.15
		fmt.Println(value)
	} else if args1 == "-K" && args2 == "-C" {
		value = value - 273.15
		fmt.Println(value)
	}
}
