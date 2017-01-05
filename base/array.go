package main

import "fmt"

func main() {

  s := make([]string, 3)
  fmt.Println("emp:", s)

  s[0] = "星"
  s[1] = "期"
  s[2] = "八"
  fmt.Println("set:", s)

  fmt.Println("get:", s[2])

  s = append(s, "不")
  s = append(s, "上", "班")
  fmt.Println("apd:", s)

  c := make([]string, len(s))
  copy(c, s)
  fmt.Println("cpy:", c)

  l := s[3:6]
  fmt.Println("sl1:", l)

  l = s[:3]
  fmt.Println("sl2:", l)

  l = s[3:]
  fmt.Println("sl3:", l)

  t := []string{"不", "上", "班"}
  fmt.Println("dcl:", t)

}
