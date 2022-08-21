package tests

import (
	"testing"

	"github.com/alfiankan/sherlock-struct-autowire/example"
	"github.com/alfiankan/sherlock-struct-autowire/sherlock"
)

var config = example.Pertalite{
	Barrel: 5000,
}
var customBody = example.Body{}

func TestWireGenFile(t *testing.T) {
	t.Run("Generating Wire file", func(t *testing.T) {

		var totalBarel int64 = 5000

		sr := sherlock.New()
		//sr.SetPath("../")
		sr.SetPkgName("tests")
		sr.Add(example.Car{}, config, customBody)
		sr.Add(example.Engine{}, totalBarel)
		sr.Add(example.Body{})
		sr.Gen()

	})
}
