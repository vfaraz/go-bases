package main

import (
	"errors"
	"fmt"
)

const (
	dog          = "dog"
	cat          = "cat"
	hamster      = "hamster"
	spider       = "spider"
	kgPerDog     = 10
	kgPerCat     = 5
	kgPerHamster = 0.250
	kgPerSpider  = 0.150
)

type funcFloat64 func(float64) float64

func main() {
	animalDog, msg := animal(dog)
	animalCat, msg := animal(cat)
	animalHamster, msg := animal(hamster)
	animalSpider, msg := animal(spider)
	if msg != nil {
		fmt.Println(msg)
	} else {
		kgDog := animalDog(10)
		kgCat := animalCat(5)
		kgHamster := animalHamster(5)
		kgSppider := animalSpider(5)

		fmt.Println("Kg necesarios para perros:", kgDog)
		fmt.Println("Kg necesarios para gatos:", kgCat)
		fmt.Println("Kg necesarios para hamsters:", kgHamster)
		fmt.Println("Kg necesarios para arañas:", kgSppider)

	}

}

func animal(animalType string) (function funcFloat64, err error) {
	switch animalType {
	case dog:
		return dogFunc, err
	case cat:
		return catFunc, err
	case hamster:
		return hamsterFunc, err
	case spider:
		return spiderFunc, err
	default:
		err = errors.New("No existe el animal")
		return
	}
}

func dogFunc(numDog float64) (kgDogFood float64) {
	kgDogFood = numDog * kgPerDog
	return
}
func catFunc(numCat float64) (kgCatFood float64) {
	kgCatFood = numCat * kgPerCat
	return
}
func hamsterFunc(numHamster float64) (kgHamsterFood float64) {
	kgHamsterFood = numHamster * kgPerHamster
	return
}
func spiderFunc(numSpider float64) (kgSpiderFood float64) {
	kgSpiderFood = numSpider * kgPerSpider
	return
}
