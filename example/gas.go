package example

import "fmt"

type IGas interface {
	Fill() error
}

type Gas struct {
	GasOil Pertalite
}

var _ IGas = &Gas{}

func (gas *Gas) Fill() error {
	fmt.Println("Fuelling gas")
	return nil
}
