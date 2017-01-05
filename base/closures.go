package main

import "fmt"

func plusLife() func() int {
  i := 0
  return func() int {
    i += 1
    return i
  }
}

func main() {

  ha := plusLife()

  fmt.Println("续命成功，目前为+", ha(), "s")
  fmt.Println("续命成功，目前为+", ha(), "s")
  fmt.Println("续命成功，目前为+", ha(), "s")
}
