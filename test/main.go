package main

import (
    "fmt"
    "io/ioutil"
)
func init() {
    data, _ := ioutil.ReadFile("./flag.txt")
    fmt.Printf("Flag: %s", string(data)) // Properly formatted string
}
func hello() {
    data, _ := ioutil.ReadFile("/test/flag.txt");
    fmt.Printf("Flag: %s", string(data)); // Properly formatted string
}
func main(){
	hello();
}