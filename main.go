package main

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

	Go.PrintClasses()
}
