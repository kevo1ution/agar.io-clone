package game

import (
	"fmt"
	"image/color"

	"github.com/google/uuid"
)

type event interface {
	OnCollide(plr Player)
}

type item struct {
	Id       uuid.UUID
	Position Point
	Velocity Point
}

func makeItem(pos Point, vel Point) item {
	return item{
		Id:       uuid.New(),
		Position: pos,
		Velocity: vel,
	}
}

type food struct {
	item
	Radius float64
	Mass   int
	Color  color.Alpha
}

func MakeFood(pos Point, vel Point, rad float64, mass int, col color.Alpha) food {
	return food{
		item:   makeItem(pos, vel),
		Radius: rad,
		Color:  col,
	}
}

func (food food) OnCollide(plr Player) {
	fmt.Println("coming from food: ", food.Id)
}

type virus struct {
	item
	Radius      float64
	Mass        int
	Fill        color.Alpha
	Stroke      color.Alpha
	StrokeWidth float64
}

func MakeVirus(pos Point, vel Point, rad float64, mass int, fill color.Alpha, stroke color.Alpha, strokeWidth float64) virus {
	return virus{
		item:        makeItem(pos, vel),
		Radius:      rad,
		Fill:        fill,
		Stroke:      stroke,
		StrokeWidth: strokeWidth,
	}
}

func (virus virus) OnCollide(plr Player) {
	fmt.Println("coming from virus: ", virus.Id)
}
