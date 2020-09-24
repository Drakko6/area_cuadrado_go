package main

import "fmt"

func main(){

// 	- Calcular el área del cuadrado
	var lado float64

	fmt.Print("Lado del cuadrado: ")
	fmt.Scan(&lado)

	area := lado*lado

	//fmt.Println("El lado es:",lado, "cm" )
	fmt.Println("El área es:", area, "cm2")


}