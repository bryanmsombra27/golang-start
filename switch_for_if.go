package main

import "fmt"

func switch_for_if() {
	should_buy := true
	debts := true

	if should_buy && !debts {
		fmt.Println("Compralo padrino")
	}	else {
		fmt.Println("No te alcanza pobre")
	}

	// variables sombra (scope)
	x:=10
	if x > 5{
		fmt.Println(x)
		x:=5
		fmt.Println(x)
	} 

	fmt.Println(x)

	// ciclo FOR
	for i:= 0; i <= 10;  i++ {
		fmt.Println(i * 2)	

	}
	n:=0

	// break
	for {
		n++
		fmt.Println(n)

		if n == 100{
			break
		}
	}

	// for range
	evenNumbers:= []int{2,4,6,8,10,12}

	// indice y valor del slice
	for i,v:= range evenNumbers{
		fmt.Println(i,v)
	}
	// solo valores
	for _,v:= range evenNumbers{
		fmt.Println(v)
	}


	// iterando mapas,m los mapas pueden regresar un orden de llaves diferentesm no devuelven posiciones fijas como los slices
	iterate_map :=map[string]int {
		"a":1,
		"b":2,
		"c":3,
	}

	for k,v:= range iterate_map{

		fmt.Println(k,v)

	}

	// iterando string
	examples:=[]string{"jhola","carambola","bola","polola"}

	// for label, colocar un alias a un ciclo
outer:
	for _,example :=range examples{
		for i,v := range example {
			fmt.Println(i,v,string(v))
			if string(v) == "l" {
				continue outer
			}
		}

	}



	//swtich
	loop:
		for i:=0; i<=10; i++ {
			switch i {
			case 0,2,4,6,8:
				fmt.Println(i, "Es Par")
				
			case 3:
				fmt.Println(i, "Es impar")
			case 7:
				fmt.Println(i, "llega la 7")
				break loop
			default:
				fmt.Println(i, "valor por defecto")
			}
		}


}
