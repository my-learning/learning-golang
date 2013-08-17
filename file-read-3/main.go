package main

import (
  "io/ioutil"
)

func main() {
  b, err := ioutil.ReadFile("sample.txt")
  if err != nil { panic(err) }

  err = ioutil.WriteFile("output.txt", b, 0644)
  if err != nil { panic(err) }
}
