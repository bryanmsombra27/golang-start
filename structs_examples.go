package main

func structs() {
	// structs
	type person struct {
		name string
		age  int
		pet  string
	}

	var ricardo person = person{
		name: "koso",
		age:  25,
		pet:  "kaso",
	}

	pedro := person{
		"koso",
		25,
		"kaso",
	}

	// struct anonimo
	var personAnonim struct {
		name string
		age  int
		pet  string
	}

	personAnonimShort := struct {
		name string
		age  int
	}{
		name: "dante",
		age:  12,
	}

	println(personAnonim)
	println(personAnonimShort)

	println(ricardo)
	// se accede/asignan los valores como un objeto literal de js
	println(ricardo.name)
	println(pedro)

	// structs anonimos

}