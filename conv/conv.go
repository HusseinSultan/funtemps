package conv

/*
  In this package, all conversion functions should be implemented.
    FahrenheitToCelsius
    CelsiusToFahrenheit
    KelvinToFahrenheit
    FahrenheitToKelvin
    CelsiusToKelvin
    KelvinToCelsius
*/

// FahrenheitToCelsius converts a temperature in degrees Fahrenheit to degrees Celsius.
func FahrenheitToCelsius(f float64) float64 {
        return (f - 32) * 5 / 9
}

// CelsiusToFahrenheit converts a temperature in degrees Celsius to degrees Fahrenheit.
func CelsiusToFahrenheit(c float64) float64 {
        return c*9/5 + 32
}

// KelvinToFahrenheit converts a temperature in Kelvin to degrees Fahrenheit.
func KelvinToFahrenheit(k float64) float64 {
        return (k-273.15)*9/5 + 32
}

// FahrenheitToKelvin converts a temperature in degrees Fahrenheit to Kelvin.
func FahrenheitToKelvin(f float64) float64 {
        return (f-32)*(5.0/9.0) + 273.15
}

// CelsiusToKelvin converts a temperature in degrees Celsius to Kelvin.
func CelsiusToKelvin(c float64) float64 {
        return c + 273.15
}

// KelvinToCelsius converts a temperature in Kelvin to degrees Celsius.
func KelvinToCelsius(k float64) float64 {
        return k - 273.15
}
