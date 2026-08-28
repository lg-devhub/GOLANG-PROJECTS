package main
import (
	"fmt"
)

func main(){
	fmt.Println("Hello World!")

	var x [5]float64

	x[0] = 3.5
	x[1] = 4.7
	x[2] = 6.9
	x[3] = 6
	x[4] = 4.6

	var total float64 = 0
	for i:=0; i < len(x); i++{
		total += x[i]
	}

	fmt.Println(total / float64(len(x)))
}