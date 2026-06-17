package main

import (
	"fmt"
)


func conversion(kelvin float64) float64 {
	return kelvin - 273.15
}


func main(){
	tempKelvin := 300.0
	tempCelsius := conversion(tempKelvin)

	fmt.Printf("%2.fK equivalem a %2.f°C\n", tempKelvin, tempCelsius)

}