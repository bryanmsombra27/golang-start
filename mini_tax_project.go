package main

import "fmt"

func taxCalculator() {	
var	(
	client_name string 
	email string
	subtotal float64
	iva float64 = 0.16
	) 


	println("Ingresa tu nombre: ")
	fmt.Scan(&client_name)
	println("Ingresa tu email: ")
	fmt.Scan(&email)
	println("Ingresa subtotal: ")
	fmt.Scan(&subtotal)

	total:= subtotal + (subtotal *iva) 


	fmt.Printf("Hola %s, el monto a pagar es de  $%.2f , tu factura se enviara al correo %s. Ten un buen dia maricon",client_name,total,email)
	



}
