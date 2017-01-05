package main

import "fmt"

func main() {

  f := make(map[string]string)

  f["萌节"] = "10/10"
  f["萝莉节"] = "10/11"
  f["光棍节"] = "11/11"
  fmt.Println("festival:", f)

  v1 := f["萌节"]
  fmt.Println("v1: ", v1)

  fmt.Println("len: ", len(f))

  delete(f, "光棍节")
  fmt.Println("festival:", f)

  _, prs := f["光棍节"]
  fmt.Println("prs:", prs)

  n := map[string]string{"扒衣见君节":"8/1"}
  fmt.Println("map:", n)
}
