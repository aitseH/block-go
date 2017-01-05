package main

import "fmt"


func yuyuko_feed(foods ...string) {
  for _, food := range foods {
    fmt.Println("yuyuko are eating", food)
  }
}

func main(){
  yuyuko_feed("dumpling", "lamian", "sushi")

  foods := []string{"tempura", "takoyaki", "curry", "stateliness"}
  yuyuko_feed(foods...)
}
