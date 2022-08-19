//go:build wireinject
// +build wireinject

package main

import (
	"bifrost-di/example"

	"github.com/google/wire"
)

func InitializeCar(gas example.Pertalite, body example.Body) example.Car {
	wire.Build(
		wire.Struct(new(example.Engine), "*"),
		wire.Struct(new(example.Gas), "*"),
		wire.Struct(new(example.Car), "*"),
		wire.Struct(new(example.Oil), "*"),
	)

	return example.Car{}
}
