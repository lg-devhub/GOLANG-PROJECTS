package main

import (
	"fmt"
	"math"
)

func somar(a, b int) int {
	return a + b
}

func subtrair(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

func dividir(a, b int) int {
	return a / b
}

//retornos múltiplos
func divisaoRestante (a, b float64) (float64, float64) {
	resultado := a / b
	resto := math.Mod(a, b)

	return resultado, resto
}

func sumValues(numbers ...float64) float64 {
	total := 0.0

	for _, number := range numbers {
		total += number
	}
	return float64(total)
}

func main() {
	somatoria := somar(80,20)
	menos := subtrair(100, 50)
	multiplica := multiplicar(100, 50)
	//divisao := dividir(100, 50)
	divisao, resto := divisaoRestante(30, 7)

	fmt.Println("Hello World!")
	fmt.Println("Soma: ", somatoria)
	fmt.Println("Subtração: ", menos)
	fmt.Println("Multiplicação: ", multiplica)
	fmt.Println(" | Divisão:", divisao, " | Restante:",  resto)

	fmt.Println(sumValues(2.5, 9.7, 3.6, 10.5, 21.45))


	//variáveis como funções

	somarrr := func(a, b int) int {
		return a + b
	}

	result := somarrr(10, 50)

	fmt.Println(result)

	
}