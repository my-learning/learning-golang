// I see at https://gist.github.com/border/775526

package main

import (
  "fmt"
  "os"
  "encoding/json"
  "io/ioutil"
)

type jsonobject struct {
  Object ObjectType 
}

type ObjectType struct {
  Buffer_size int
  Databases   []DatabaseType
}

type DatabaseType struct {
  Host    string
  User    string
  Pass    string
  Type    string
  Name    string
  Tables  []TablesType
}

type TablesType struct {
  Name      string
  Statement string
  Regex     string
  Types     []TypesType
}

type TypesType struct {
  Id    string
  Value string
}

func main() {
  file, e := ioutil.ReadFile("./setting.json")
  if e != nil {
    fmt.Printf("File error: %v\n", e)
    os.Exit(1)
  }
  fmt.Printf("%s\n", string(file))

  var jsontype jsonobject
  json.Unmarshal(file, &jsontype)
  fmt.Printf("Results: %v\n", jsontype)
}
