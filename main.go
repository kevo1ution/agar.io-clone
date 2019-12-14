package main

import (
	"sync"
	"time"

	"github.com/kevo1ution/agar.io-clone/config"
)

func main() {
	config.Setup()
	/*
		//Item manager example
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
	*/

	//start loops
	var wg sync.WaitGroup
	timedLoop(&wg, gameLoop, 100*time.Millisecond)
	timedLoop(&wg, moveLoop, 100*time.Millisecond)
	timedLoop(&wg, sendUpdates, 100*time.Millisecond)
	wg.Wait()
}

func timedLoop(wg *sync.WaitGroup, fn func() bool, interval time.Duration) {
	wg.Add(1)
	go func() {
		timer := time.NewTicker(interval)
		for range timer.C {
			if fn() {
				break
			}
		}
		wg.Done()
	}()
}

// every one seconds
func gameLoop() bool {

	return false
}

func moveLoop() bool {

	return false
}

func sendUpdates() bool {

	return false
}
