package main

import "fmt"

func main() {
	Go := Course{
		Name:    "Go desde cero",
		Price:   12.34,
		IsFree:  false,
		UserIds: []uint{12, 56, 89},
		Classes: map[uint]string{
			1: "Introducción",
			2: "Estructuras",
			3: "Maps",
		},
	}

	React := Course{
		Name:   "React desde cero",
		IsFree: true,
	}

	fmt.Println(Go.Name)
	fmt.Println(React.Classes)
}
