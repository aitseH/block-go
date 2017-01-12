package main

import "fmt"

type person struct {
	name string
	age  string
}

func main() {

	fmt.Println(person{"Remu", "16 - 24"})

	fmt.Println(&person{"marisa", "16 - 24"})

	fmt.Println(person{"Yukari", "5oiHJN(*&9*&U(*%$#d_()*)(.,"})

	fmt.Println(person{"Remilia", "500++"})

	fmt.Println(person{"Furandōru", "495++"})

	s := person{name: "Alice", age: "16 - 24"}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = "16"
	fmt.Println(sp.age)
}
