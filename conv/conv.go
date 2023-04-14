package conv

/*
  I denne pakken skal alle konverteringfunksjonene
  implementeres. Bruk engelsk.
    FarhenheitToCelsius
    CelsiusToFahrenheit
    KelvinToFarhenheit
    ...
*/

// FarhenheitToCelsius converts a temperature in degrees Fahrenheit to degrees Celsius.
func FarhenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

// CelsiusToFahrenheit converts a temperature in degrees Celsius to degrees Fahrenheit.
func CelsiusToFahrenheit(c float64) float64 {
	return c*9/5 + 32
}

// KelvinToFarhenheit converts a temperature in Kelvin to degrees Fahrenheit.
func KelvinToFarhenheit(k float64) float64 {
	return (k-273.15)*9/5 + 32
}
func FahrenheitToKelvin(f float64) float64 {
	return (f-32)*(5.0/9.0) + 273.15
}
func CelsiusToKelvin(C float64) float64 {
	return (C + 273.15)
}
func KelinToCelsius(K float64) float64 {
	return (K - 273.15)
}
