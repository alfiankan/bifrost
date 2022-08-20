package sherlock

const (
	wireHead = `
//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"fmt"
`
	wireInit = `
)

func init() {
	fmt.Println("Initializer")
}

`
	initializerTemplate = `
func Initialize%s(%s) %s {
	wire.Build(
		%s
	)
	return %s{}
}
`
)
