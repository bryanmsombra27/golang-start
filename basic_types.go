package main

func basic_types() {
	var a int = 10
	var b float32 = 3.14
	var c string = "hola"
	var d bool = true
	var e rune = 'ñ'
	var f complex128 = 2 + 3i

	// CAPACIDAD DE NUMEROS
	var g int8 = 127 // rango entre  -128  y 127 numeros enteros

	var h uint8 = 127 // rango entre  0 y 255 numeros enteros y solo positivos

	println(a, b, c, d, e, f, g, h)

}
