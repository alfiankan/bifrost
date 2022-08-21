package tests

import (
	"testing"

	"github.com/alfiankan/sherlock-struct-autowire/example"
	"github.com/alfiankan/sherlock-struct-autowire/sherlock"
)

func TestWireGenFileBuilder(t *testing.T) {
	t.Run("Generating Wire file using builder patern", func(t *testing.T) {

		var config = example.Pertalite{
			Barrel: 5000,
		}
		var customBody = example.Body{}

		var totalBarel int64 = 5000

		sherlock.New().
			SetPkgName("tests").
			Add(example.Car{}, config, customBody).
			Add(example.Engine{}, totalBarel).
			Add(example.Body{}).
			Gen()
	})
}
