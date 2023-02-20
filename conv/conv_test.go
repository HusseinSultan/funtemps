package conv

import (
	"reflect"
	"testing"
)

/*
*
	Mal for testfunksjoner
	Du skal skrive alle funksjonene basert på denne malen
	For alle konverteringsfunksjonene (tilsammen 6)
	kan du bruke malen som den er (du må selvsagt endre
	funksjonsnavn og testverdier)
*/
func TestFarhenheitToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 134, want: 56.67},
	}

	for _, tc := range tests {
		got := FarhenheitToCelsius(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

// De andre testfunksjonene implementeres her
// ...
package conv

import (
	"reflect"
	"testing"
)

func TestFarhenheitToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 134, want: 56.67},
		{input: 32, want: 0},
		{input: -40, want: -40},
	}

	for _, tc := range tests {
		got := FarhenheitToCelsius(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestCelsiusToFarhenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 56.67, want: 134},
		{input: 0, want: 32},
		{input: -40, want: -40},
	}

	for _, tc := range tests {
		got := CelsiusToFarhenheit(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestMetersToFeet(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 1, want: 3.28084},
		{input: 0, want: 0},
		{input: -10, want: -32.8084},
	}

	for _, tc := range tests {
		got := MetersToFeet(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestFeetToMeters(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 3.28084, want: 1},
		{input: 0, want: 0},
		{input: -32.8084, want: -10},
	}

	for _, tc := range tests {
		got := FeetToMeters(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestKilogramsToPounds(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 1, want: 2.20462},
		{input: 0, want: 0},
		{input: -5, want: -11.0231},
	}

	for _, tc := range tests {
		got := KilogramsToPounds(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestPoundsToKilograms(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 2.20462, want:
			package conv

			import (
				"reflect"
				"testing"
			)
			
			func TestPoundsToKilograms(t *testing.T) {
				type test struct {
					input float64
					want  float64
				}
			
				tests := []test{
					{input: 2.20462, want: 1},
					{input: 0, want: 0},
					{input: -11.0231, want: -5},
				}
			
				for _, tc := range tests {
					got := PoundsToKilograms(tc.input)
					if !reflect.DeepEqual(tc.want, got) {
						t.Errorf("expected: %v, got: %v", tc.want, got)
					}
				}
			}
			
			func TestGallonsToLiters(t *testing.T) {
				type test struct {
					input float64
					want  float64
				}
			
				tests := []test{
					{input: 1, want: 3.78541},
					{input: 0, want: 0},
					{input: -5, want: -18.9271},
				}
			
				for _, tc := range tests {
					got := GallonsToLiters(tc.input)
					if !reflect.DeepEqual(tc.want, got) {
						t.Errorf("expected: %v, got: %v", tc.want, got)
					}
				}
			}
			
			func TestLitersToGallons(t *testing.T) {
				type test struct {
					input float64
					want  float64
				}
			
				tests := []test{
					{input: 3.78541, want: 1},
					{input: 0, want: 0},
					{input: -18.9271, want: -5},
				}
			
				for _, tc := range tests {
					got := LitersToGallons(tc.input)
					if !reflect.DeepEqual(tc.want, got) {
						t.Errorf("expected: %v, got: %v", tc.want, got)
					}
				}
			}
			