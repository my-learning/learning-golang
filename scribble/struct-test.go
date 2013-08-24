package main

import (
  "fmt"
)

type Vertex struct {
  x int
  y int
}

func (v *Vertex) SetX(x int) {
  v.x = x
}

func (v Vertex) GetX() int {
  return v.x
}

func (v *Vertex) SetY(y int) {
  v.y = y
}
func (v Vertex) GetY() int {
  return v.y
}

func main() {
  v := new(Vertex)
  v.SetX(100)
  v.SetY(200)

  fmt.Printf("v.X is %d\n", v.GetX())
  fmt.Printf("v.Y is %d\n", v.GetY())

  ary := [3]int {1, 2, 3}
  fmt.Println(len(ary))

  no_named_struct_ary := []struct { i int }{ {1}, {2}, {3} }
  fmt.Println(no_named_struct_ary[0])

  slice := make([]int, 0, 5)
  fmt.Println(len(slice), cap(slice))

  m := map[string]int {
    "abc": 123,
    "def": 456,
  }

  fmt.Println("m[\"abc\"] is %d", m["abc"])
}
