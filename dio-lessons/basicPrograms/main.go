package main

import (
	"fmt"
)

const Fahrenheit float64 = 212.0

func main() {
	tempF := Fahrenheit
	tempC := (Fahrenheit - 32) * 5/9

	fmt.Println("A temperatura em Fahrenheit é:",tempF)
	fmt.Println("A temperatura convertida em Celsius é:",tempC)
	//ou
	fmt.Printf("A temperatura em Fahrenheit é: %g e em Celsius é: %g ", tempF, tempC)
}