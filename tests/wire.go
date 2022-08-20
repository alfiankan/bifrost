//go:build wireinject
// +build wireinject

package tests

import (
	"bifrost-di/example"
	"fmt"
	"github.com/google/wire"
)

func init() {
	fmt.Println("Initializer")
}

func InitializeCar(pertalite example.Pertalite, body example.Body) example.Car {
	wire.Build(
		wire.Struct(new(example.Car), "*"),
		wire.Struct(new(example.Engine), "*"),
		wire.Struct(new(example.Gas), "*"),
		wire.Struct(new(example.Oil), "*"),
	)
	return example.Car{}
}

func InitializeEngine(pertalite example.Pertalite) example.Engine {
	wire.Build(
		wire.Struct(new(example.Engine), "*"),
		wire.Struct(new(example.Gas), "*"),
		wire.Struct(new(example.Oil), "*"),
	)
	return example.Engine{}
}

func InitializeBody() example.Body {
	wire.Build(
		wire.Struct(new(example.Body), "*"),
	)
	return example.Body{}
}
