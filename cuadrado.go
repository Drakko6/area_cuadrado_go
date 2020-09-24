package main

import "fmt"


func suma(a int, b int) int {
	
	return a + b
}

func resta(a int, b int) int {
	
	return a - b
}





func main(){

// 	- Calcular el área del cuadrado
	var lado float64

	fmt.Print("Lado del cuadrado: ")
	fmt.Scan(&lado)

	area := lado*lado

	fmt.Println("El lado es:",lado, "cm" )
	fmt.Println("El área es:", area, "cm2")


	a := suma(2, 3)
	fmt.Println("La suma de numeros es:", a )

}