package game

type Point struct {
	X float64
	Y float64
}

type Character struct {
	Position Point
	Velocity Point
	radius   float64
}

type Player struct {
	Char  Character
	Score int
}
