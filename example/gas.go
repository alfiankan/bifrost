package example

import "fmt"

type IGas interface {
	Fill() error
}

type Gas struct {
	GasOil Pertalite
}

func (gas *Gas) Fill() error {
	fmt.Println("Fuelling gas")
	return nil
}
