package animal

import "fmt"

type AnimalDetailStruct struct {
	Name           string
	Color          string
	Speed          int
	IsPoisones     bool
	Carnivores     bool
	goingtoExtinct bool
}

func PrintAnimalName(animal AnimalDetailStruct) {
	fmt.Println(animal.Name)
}
