package main

import "fmt"

func main() {

	messages := make(chan string, 2)

	messages <- "mikusama"
	messages <- "saiko!!!"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
