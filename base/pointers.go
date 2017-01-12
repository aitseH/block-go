package main

import "fmt"

func playGame(name string) {
	name = "ysatnaf"
}

func changeName(name *string) {
	*name = "nanoo"
}

func main() {
	name := "ysatnaf"

	fmt.Println("昵称:", name)

	playGame(name)
	fmt.Println("游戏昵称:", name)

	changeName(&name)
	fmt.Println("改名后:", name)

	fmt.Println("你家住哪:", &name)
}
