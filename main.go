package main

import (
	"bifrost-di/bifrost"
	"bifrost-di/config"
	"bifrost-di/example"
	"fmt"
)

func NewCar() example.Car {
	return example.Car{}
}

func NewEngine() example.Engine {
	return example.Engine{}
}

func NewBody() example.Body {
	return example.Body{}
}

func main() {

	// manual dependency

	// =================================================== //
	//						Prototype					   //
	// =================================================== //
	fmt.Println()
	fmt.Println("========================================")
	fmt.Println()

	rootBf := bifrost.Bifrost{}

	rootBf.Gen(example.Car{}, config.Config2)

}
