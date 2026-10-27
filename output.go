package main

import "fmt"

func inputOutput() {
var age int
var n string


fmt.Println("Ingresa tu nombre: ")
fmt.Scan(&n)

fmt.Println("Ingresa tu edad: ")
fmt.Scan(&age)

fmt.Printf("Tu nombre es %s y tienes %d",n,age)

}
