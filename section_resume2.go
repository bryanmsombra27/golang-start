package main

import "fmt"

//  structs
	type Item struct {
		sku      string
		name     string
		price    int
		quantity int
	}
	type Order struct {
		id       string
		customer string
		items    []Item
		quantity int
		meta     map[string]string
	}

func section_resume() {

	producto := Item{
		sku:      "paoq12",
		name:     "producto 1",
		price:    12,
		quantity: 1243,
	}

	orden := Order{
		id:       "1",
		customer: "aba",
		items: []Item{
			producto,
			producto,
			producto,
		},
		quantity: 3,
		meta: map[string]string{
			"city": "Acapulco",
			"source": "map",
		},
	}

	fmt.Println(orden.id)
	fmt.Println(orden.customer)
	fmt.Println(orden.items[0])
	fmt.Println(orden.meta["cupon"])
	orden.meta["cupon"] =  "HOLACARADEBOLA"
	fmt.Println(orden.meta["cupon"])
}