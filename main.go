package main

import (
	"fmt"
	"image/color"

	"github.com/kevo1ution/agar.io-clone/config"
	"github.com/kevo1ution/agar.io-clone/game"
)

func main() {
	config.Setup()
	itemManager := game.MakeItemManager()
	fmt.Printf("%v", itemManager)

	food := game.MakeFood(game.Point{0, 0}, game.Point{0, 0}, 2.0, 2, color.Alpha{})
	virus := game.MakeVirus(game.Point{0, 0}, game.Point{0, 0}, 2.0, 2, color.Alpha{}, color.Alpha{}, 1.0)
	itemManager.Add(food.Id, food)
	itemManager.Add(virus.Id, virus)
	fmt.Printf("%v\n\n", itemManager)

	//iterate through the map and trigger polymorphism
	for _, v := range itemManager.Items {
		v.OnCollide(game.Player{})
	}
}
