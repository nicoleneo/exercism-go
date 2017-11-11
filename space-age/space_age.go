package space

import "math"

type Planet string

var earthYearsMap = map[Planet]float64{
	"Earth":   1,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

func Age(ageInSeconds float64, planet Planet) float64 {
	earthYearInSeconds := 31557600.0
	ageInYears := ageInSeconds / earthYearsMap[planet] / earthYearInSeconds
	return toFixed(ageInYears, 2)
}

// https://stackoverflow.com/a/29786394
func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
