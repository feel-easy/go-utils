package main

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

type Animal struct {
	Name  string `json:"name"`
	Order string `json:"order"`
}

func main() {
	var animals []Animal
	var animals1 []Animal
	animals = append(animals, Animal{Name: "Platypus", Order: "Monotremata"})
	animals = append(animals, Animal{Name: "Quoll", Order: "Dasyuromorphia"})

	jsonStr, err := jsoniter.Marshal(&animals)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(string(jsonStr))
	jsoniter.Unmarshal(jsonStr, &animals1)
	fmt.Printf("animals1 is %v", animals1)
}
