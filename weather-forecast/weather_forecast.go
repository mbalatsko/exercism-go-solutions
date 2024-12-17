// Package weather forecast the current weather condition of various cities in Goblinocus.
package weather

// CurrentCondition represents a current weather condition.
var CurrentCondition string
// CurrentLocation represents current location.
var CurrentLocation string

// Forecast returns an string value which represent forecast for cite and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
