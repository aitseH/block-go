package main

func rec() {
  print("+1s")
  if true {
    rec()
  }
}

func main() {
  rec()
}
