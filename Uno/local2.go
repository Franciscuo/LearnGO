package main

import "fmt"

func main() {
	var nota1, nota2, nota3 int
	fmt.Print("Ingrese primer nota:")
	fmt.Scan(&nota1)
	fmt.Print("Ingrese segundo nota:")
	fmt.Scan(&nota2)
	fmt.Print("Ingrese tercer nota:")
	fmt.Scan(&nota3)
	promedio := (float32(nota1) + float32(nota2) + float32(nota3)) / 3
	fmt.Println("Promedio:", promedio)
	if promedio >= 7 {
		fmt.Print("Promocionado")
	}
}
