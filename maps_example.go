package main

import (
	"fmt"
	"maps"
)

func maps_example() {
var nilMap map[string]int  
totalWins := map[string]int{}
// asignacion de valor a un map
totalWins["lsp"] =2


fmt.Printf("nilMap: %v\n", nilMap)
fmt.Println(totalWins)
//  map literal :inicializando un map con valores por defecto
teams:= map[string][]string{
"ironman":[]string{
	"widow","war machine",
},
"captain":[]string{
	"antman","winter soldier",
},
}

println(teams)
println(teams["ironman"][0])

// TODOS LOS QUE NO QUE PUEDAN COMPRAR  CON == O != NO  PUEDEN SER LLAVES
	age:= make(map[int][]string,10)
	fmt.Println(age)

	my_map:= map[string]int {
		"hola":1,
		"mundo":2,
	}
	my_map_b:= map[string]int {
		"mundo":2,
		"hola":1,
	}
	// comma, ok idiom: [se asemeja a una destructuracion de objeto donde obtemeos el valor  y un booleno que indica si la llave existe o no]
	value,ok:= my_map["hola"]

	fmt.Println(value,ok)


	// eliminar de un map un elemento
	delete(my_map,"hola")

	fmt.Println(my_map)

	// dejar el mapa vacio limpiar todos los elementos
	 clear(my_map)

	// comparando maps para saber si son iguales sin importar el orden de las llaves solo que el contenido sea igual
	fmt.Println(maps.Equal(my_map,my_map_b))
	

	}
