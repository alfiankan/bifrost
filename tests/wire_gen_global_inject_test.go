package tests

import (
	"testing"

	"github.com/alfiankan/sherlock-struct-autowire/example"
	"github.com/alfiankan/sherlock-struct-autowire/sherlock"
)

func TestWireGenFileBuilderGlobalInject(t *testing.T) {
	t.Run("Generating Wire file using builder patern global inject", func(t *testing.T) {

		var config = example.Pertalite{
			Barrel: 5000,
		}
		var customBody = example.Body{}

		sherlock.New().
			SetPkgName("tests").
			Add(example.Car{}, customBody).
			Add(example.Engine{}).
			Add(example.Body{}).
			SetGlobalInject(config).
			Gen()
	})
}
