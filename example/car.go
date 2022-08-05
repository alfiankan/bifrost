package example

type ICar interface {
	Get() (string, error)
}

type Car struct {
	Engine Engine
	Body   Body
}

func (car *Car) Get() (string, error) {
	car.Engine.Start()
	car.Body.Close()
	return "you got the car", nil
}
