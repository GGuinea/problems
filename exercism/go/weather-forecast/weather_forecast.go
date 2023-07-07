// Package weather Represents a package that provides weather forecast functionality.
package weather

// CurrentCondition  a variable that stores the current weather condition.
var CurrentCondition string
// CurrentLocation a variable that stores the current location.
var CurrentLocation string

// Forecast function is responsible for providing the weather forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
