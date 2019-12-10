// Package stringutil contains utility functions for working with strings.
package calcutil

import (
	"math"
	"math/rand"
	"regexp"

	"github.com/kevo1ution/agar.io-clone/config"
)

//validate a nick name
func ValidNick(nickname string) bool {
	var isValidNick = regexp.MustCompile(`^\w*$`)

	return isValidNick.MatchString(nickname)
}

func MassToRadius(mass float64) float64 {
	return 4 + math.Sqrt(mass)*6
}

func GetDistance(p1 Circle, p2 Circle) float64 {
	return math.Sqrt(math.Pow(p2.Center.X-p1.Center.X, 2)+math.Pow(p2.Center.Y-p1.Center.Y, 2)) - p1.Radius - p2.Radius
}

func RandomInRange(from float64, to float64) float64 {
	return math.Floor(rand.Float64()*(to-from)) + from
}

func RandomPosition(radius float64) Point {
	return Point{
		X: RandomInRange(radius, config.Config.Width-radius),
		Y: RandomInRange(radius, config.Config.Height-radius),
	}
}

func UniformPosition(circles []Circle, radius float64) Point {
	var bestCandidate Circle
	maxDistance := 0.0
	numberOfCandidates := 10

	if len(circles) == 0 {
		return RandomPosition(radius)
	}

	// pick the furthest position
	for ci := 0; ci < numberOfCandidates; ci++ {
		// Getting a random position, and picking the closest
		// to a candidate
		minDistance := math.Inf(1)
		candidate := Circle{
			Center: RandomPosition(radius),
			Radius: radius,
		}

		for pi := 0; pi < len(circles); pi++ {
			distance := GetDistance(candidate, circles[pi])
			if distance < minDistance {
				minDistance = distance
			}
		}

		if minDistance > maxDistance {
			bestCandidate = candidate
			maxDistance = minDistance
		} else {
			return RandomPosition(radius)
		}
	}

	return bestCandidate.Center
}
