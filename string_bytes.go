package main

import "fmt"

func string_bytes() {
	var chain string = "hola mundo"
	var my_byte  byte = chain[0]
	 s2 := chain[3:7]
	fmt.Println(chain[0])  // valor de utf-8 
	fmt.Println(my_byte)  // valor de utf-8 
	fmt.Println(s2)
	
		// string a slice
	var s string = "mother fucker"
	var bs []byte = []byte(s)
	var rs []rune = []rune(s)
	fmt.Println(bs)
	fmt.Println(rs)
}
