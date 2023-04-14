package main

import (
	"flag"
	"fmt"

	"github.com/HusseinSultan/funtemps/conv"
)

// Definerer flag-variablene i hoved-"scope"
var fahr float64
var out string
var funfacts string
var kelvin float64
var celsius float64

func main() {
	// Define command line flags

	flag.Float64Var(&fahr, "F", 0.0, "temperature in degrees Fahrenheit")
	flag.Float64Var(&kelvin, "K", 0.0, "temperature in degrees Fahrenheit")
	flag.Float64Var(&celsius, "C", 0.0, "temperature in degrees Fahrenheit")
	flag.StringVar(&out, "out", "celsius", "output temperature scale (celsius, fahrenheit, or kelvin)")
	flag.StringVar(&funfacts, "funfacts", "", "fun facts about temperature conversion")

	// Parse command line flags
	flag.Parse()

	// Convert temperature
	var result float64
	if isFlagPassed("F") {
		switch out {
		case "C":
			result = conv.FarhenheitToCelsius(fahr)
		case "F":
			result = fahr
		case "K":
			result = conv.FahrenheitToKelvin(fahr)
		default:
			fmt.Printf("Invalid output scale: %s\n", out)
			return
		}
	}
	if isFlagPassed("C") {
		switch out {
		case "F":
			result = conv.CelsiusToFahrenheit(celsius)
		case "C":
			result = celsius
		case "K":
			result = conv.CelsiusToKelvin(celsius)
		default:
			fmt.Printf("Invalid output scale: %s\n", out)
			return
		}
	}
	if isFlagPassed("K") {
		switch out {
		case "F":
			result = conv.KelvinToFarhenheit(kelvin)
		case "K":
			result = kelvin
		case "C":
			result = conv.KelinToCelsius(kelvin)
		default:
			fmt.Printf("Invalid output scale: %s\n", out)
			return
		}

	}

	// Print temperature in the requested scale
	fmt.Printf("%.2f degrees %s\n", result, out)

	// Print fun facts
	if funfacts != "" {
		fmt.Println(funfacts)
	}
}

// Bruker init (som anbefalt i dokumentasjonen) for å sikre at flagvariablene
// er initialisert.

// Funksjonen sjekker om flagget er spesifisert på kommandolinje
// Du trenger ikke å bruke den, men den kan hjelpe med logikken
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
