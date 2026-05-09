// Package weather provides a tool  that provides forecast. 
// The parameters are the current location and the condition of current location. 
// insert city and the condition. 
// you're returned a string with the current location and the weather condition.
package weather

var (
    // CurrentCondition represents what the weather currently looks like.
	CurrentCondition string 
    // CurrentLocation represents where you're currently reporting from.
	CurrentLocation  string 
)


// Forecast returns a string providing the location and the condtion of the weather there.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
