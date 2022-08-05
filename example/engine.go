package example

import "fmt"

type IEngine interface {
	Start() error
}

type Engine struct {
	Gas Gas
	Oil Oil
}

func (engine *Engine) Start() error {
	engine.Gas.Fill()
	fmt.Println("Starting the engine")
	return nil
}
