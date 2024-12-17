package meteorology

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type
func (sc TemperatureUnit) String() string {
	units := []string{"°C", "°F"}
	return units[sc]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Add a String method to the Temperature type
func (sc Temperature) String() string {
	return fmt.Sprintf("%d %s", sc.degree, sc.unit)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit
func (sc SpeedUnit) String() string {
	units := []string{"km/h", "mph"}
	return units[sc]
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed
func (sc Speed) String() string {
	return fmt.Sprintf("%d %s", sc.magnitude, sc.unit)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData
func (sc MeteorologyData) String() string {
	return fmt.Sprintf("%s: %s, Wind %s at %s, %d%% Humidity", sc.location, sc.temperature, sc.windDirection, sc.windSpeed, sc.humidity)
}
