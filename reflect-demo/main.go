package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	Name string
}

func (m *Animal) Say() {
	fmt.Println(m.Name)
}
func main() {
	animal := Animal{"张三"}

	value := reflect.ValueOf(&animal)

	f := value.MethodByName("Say")
	f.Call([]reflect.Value{})
}
