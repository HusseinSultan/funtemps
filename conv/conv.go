package conv

/*
  I denne pakken skal alle konverteringfunksjonene
  implementeres. Bruk engelsk.
    FarhenheitToCelsius
    CelsiusToFahrenheit
    KelvinToFarhenheit
    ...
*/
package conv

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

// Konverterer Farhenheit til Celsius
func FarhenheitToCelsius(value float64) float64 {
	// Her skal du implementere funksjonen
	// Du skal ikke formattere float64 i denne funksjonen
	// Gjør formattering i main.go med fmt.Printf eller
	// lignende
	return 56.67
  func FarhenheitToCelsius(value float64) float64 {
    return (value - 32) * 5 / 9
  }
  
}

// De andre konverteringsfunksjonene implementere her
// ...