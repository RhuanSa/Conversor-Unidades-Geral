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
	}
	fmt.Println(ConvertTemp(value))
}

func ConvertTemp(value float64) float64 {
	args1, args2 := os.Args[2], os.Args[3]
	if args1 == "-cl" && args2 == "-fh" {
		value = (value * 9 / 5) + 32
	} else if args1 == "-fh" && args2 == "-cl" {
		value = (value - 32) * 5 / 9
	} else if args1 == "-cl" && args2 == "-kv" {
		value = value + 273.15
	} else if args1 == "-kv" && args2 == "-cl" {
		value = value - 273.15
	}
	return value
}
