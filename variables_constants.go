package main

import (
	"fmt"
	"strconv"
)

const app = "constante no debe cambiar"
const MaxUsers int = 100
const pi float64 = 3.14159

func variables_constants(){
	// TIPADO AUTOMATICO
	var edad = 30

	var number float64 = pi

	// TIPADO MANUAL
	var name string = "kesero"
	// otra forma de declarar variables cuando el tipo es obvio
	last_name := "ochoa"
	age:=22

	fmt.Printf("hello world %s %s %s %s\n", name, last_name, strconv.Itoa(edad), strconv.Itoa(age))
	fmt.Println(app, MaxUsers, strconv.FormatFloat(number, 'f', 2, 64))
}
