package main
import (
    "fmt"
    "io/ioutil"
)
func init() {
    data, _ := ioutil.ReadFile("flag.txt")
    fmt.Printf("Flag: %s", data)
}