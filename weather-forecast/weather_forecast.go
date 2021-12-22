// Package weather  <Calculates the location's forecast>.
package weather
// 
// CurrentCondition stores the current condition.
var CurrentCondition string
// CurrentLocation stores the current Location.
var CurrentLocation string
// Forecast returns a string value equals to the Location and Weather current condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
