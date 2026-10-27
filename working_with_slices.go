package main

import (
	"fmt"
	"slices"
)

func working_with_slices() {
	// array valor fijo
	var array1 = [3]int{1, 2, 3}

	// slice valor dinamico
	var slice1 = []int{1, 2, 4, 5}

	// crear un  slice con tamaño fijo [0,0,0,0,0]
	// make_slice:= make([]int,5)
	// crear slice con tamaño fijo pero con un limite de capacidad: [] se llena al tener una longitud de 10 valores si se sobre pasa incrementa la capacidad (se defiinio en 10 asique aumenta otros 10)
	make_slice:= make([]int,0,10)

	slice1[2] = 123

	// comparando slices
	s := []string{"vio", "rem", "em"}
	x := []int{1, 23, 3, 4, 5, 6, 7}
	y := []int{1, 3, 3, 4, 5}
	z := []int{1, 3, 3, 4, 5}

	fmt.Printf("%v ",array1)
	fmt.Print(slice1)

	// comparar slice
	fmt.Print(slices.Equal(y,z))

	fmt.Print(s,x)
	fmt.Print(len(slice1))
	fmt.Println((make_slice))
	// FUNCIONES PARA SLICES
	my_slice := []int{1,2,3}

	my_slice = append(my_slice, 1)

	// vaciar slice: todos los valores se setean en 0 no desaparece su longitud
	clear(make_slice)

	// cortes en los slices
	string_slice:=[]string {"a","b","c","d","e","f"}

	slice_of_slice:= 	string_slice[0:2] // lo mismo seria declararlo asi [:2] [a,b]
	slice_cut:= 	string_slice[1:3] // [b,c]
	
	// COPIAR SLICE
	sluce_dot_cut := string_slice[:] // copia el slice (con la direccion exacta en memoria) podria sobre escribir el slice original
	copia:= make([]string,5)
	slice_copy:= copy(copia,string_slice) //copia el slice sin riesgo de sobre scribir el original 

	fmt.Println(string_slice)
	fmt.Println(slice_of_slice)
	fmt.Println(slice_cut)
	fmt.Println(sluce_dot_cut) //
	fmt.Println(copia) //
	fmt.Println(slice_copy) // 


}
