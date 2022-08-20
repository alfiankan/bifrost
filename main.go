package main

import (
	"bifrost-di/config"
	"bifrost-di/example"
	"bifrost-di/sherlock"
	"os"
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

var bodd = example.Body{}

func main() {

	sr := sherlock.New()

	sr.Add(example.Car{}, config.Config2, bodd)
	sr.Add(example.Engine{}, config.Config2)

	if len(os.Args) > 1 && os.Args[1] == "sr-gen" {
		sr.Gen()
	}

}
