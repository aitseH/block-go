package main

import "fmt"

func baka() (int) {
  return 9
}

func biggest32() (int, int) {
  return 2147483647, -2147483647
}

func main() {
  cirno := baka()
  fmt.Println(cirno)

  max, min := biggest32()
  fmt.Println(max, min)

  _, justMin := biggest32()
  fmt.Println(justMin)
}
