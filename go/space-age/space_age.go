package space

type Planet string

// Age calculates age depending on what planet the person is on
func Age(seconds float64, planet Planet) float64 {
	orbitalPeriodMap := make(map[Planet]float64)

	orbitalPeriodMap["Mercury"] = 0.2408467 * 31557600
	orbitalPeriodMap["Venus"] = 0.61519726 * 31557600
	orbitalPeriodMap["Earth"] = 1 * 31557600
	orbitalPeriodMap["Mars"] = 1.8808158 * 31557600
	orbitalPeriodMap["Jupiter"] = 11.862615 * 31557600
	orbitalPeriodMap["Saturn"] = 29.447498 * 31557600
	orbitalPeriodMap["Uranus"] = 84.016846 * 31557600
	orbitalPeriodMap["Neptune"] = 164.79132 * 31557600

	return seconds / orbitalPeriodMap[planet]
}
