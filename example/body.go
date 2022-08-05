package example

import "fmt"

type IBody interface {
	Close() error
}

type Body struct {
}

func (body *Body) Close() error {
	fmt.Println("Closing the body")
	return nil
}
