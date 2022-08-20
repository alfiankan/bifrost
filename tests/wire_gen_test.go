package tests

import (
	"bifrost-di/example"
	"bifrost-di/sherlock"
	"testing"
)

var config = example.Pertalite{
	Barrel: 5000,
}
var customBody = example.Body{}

func TestWireGenFile(t *testing.T) {
	t.Run("Generating Wire file", func(t *testing.T) {

		sr := sherlock.New()
		//sr.SetPath("../")
		sr.SetPkgName("tests")
		sr.Add(example.Car{}, config, customBody)
		sr.Add(example.Engine{}, config)
		sr.Add(example.Body{})
		sr.Gen()

	})
}
