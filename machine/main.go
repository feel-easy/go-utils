package main

import (
	"fmt"

	"github.com/feel-easy/go-utils/machine/server"
)

func main() {
	v, err := server.GetServerInfo()
	if err != nil {
		return
	}
	fmt.Printf("Os: %v, Cpu:%v, Ram:%v\n", v.Os, v.Cpu, v.Ram)
}
