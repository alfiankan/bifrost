package main

import (
	"bifrost-di/bifrost"
	"bifrost-di/example"
	"fmt"
)

func main() {

	// manual dependency

	car := example.Car{
		example.Engine{
			example.Gas{
				example.Pertalite{
					Brand:  "Shell",
					Liters: 20,
				},
			},
			example.Oil{},
		},
		example.Body{},
	}

	fmt.Println(car.Get())

	// =================================================== //
	//						Prototype					   //
	// =================================================== //
	fmt.Println()
	fmt.Println("========================================")
	fmt.Println()

	rootBf := bifrost.Bifrost{}

	rootBf.OverrideGlobal(&example.Pertalite{
		Brand:  "Endurance",
		Liters: 10,
	})

	rootBf.Get(example.Car{})

	fmt.Println()

	rootBf.Get(example.Engine{})

}
