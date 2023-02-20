package funfacts

import (
	"reflect"
	"testing"
)

/*
*
	Mal for TestGetFunFacts funksjonen.
	Definer korrekte typer for input og want,
	og sette korrekte testverdier i slice tests.
*/
func TestGetFunFacts(t *testing.T) {
	type input struct {
		temperature float64
		unit        string
	}
	type want struct {
		facts []string
		err   error
	}
	tests := []struct {
		name  string
		input input
		want  want
	}{
		{
			name: "valid temperature and unit",
			input: input{
				temperature: 0,
				unit:        "C",
			},
			want: want{
				facts: []string{"The coldest possible temperature is -273.15°C (-459.67°F)."},
				err:   nil,
			},
		},
		// Add more test cases here
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotFacts, gotErr := GetFunFacts(tc.input.temperature, tc.input.unit)
			if !reflect.DeepEqual(tc.want.facts, gotFacts) {
				t.Errorf("test failed for input %v; expected facts %v, but got %v", tc.input, tc.want.facts, gotFacts)
			}
			if !reflect.DeepEqual(tc.want.err, gotErr) {
				t.Errorf("test failed for input %v; expected error %v, but got %v", tc.input, tc.want.err, gotErr)
			}
		})
	}
}

func TestGetFunFacts(t *testing.T) {
	type test struct {
		input // her må du skrive riktig type for input
		want  // her må du skrive riktig type for returverdien
	}
	type test struct {
		input string
		want  string
	}
	func TestGetFunFacts(t *testing.T) {
		tests := []test{
			{input: "en", want: "Did you know that Norway knighted a penguin in 2008?"},
			{input: "no", want: "Visste du at det finnes en gate i Oslo som heter Dronning Mauds gate?"},
			{input: "fr", want: "Saviez-vous que les étoiles de mer n'ont pas de cerveau?"},
			{input: "de", want: "Wusstest du, dass es in Berlin mehr Brücken als in Venedig gibt?"},
			{input: "es", want: "Sabías que los osos polares son zurdos?"},
		}
	
		for _, tc := range tests {
			got := GetFunFacts(tc.input)
			if got != tc.want {
				t.Errorf("for input %q, expected: %q, but got: %q", tc.input, tc.want, got)
			}
		}
	}
		
	
	func TestGetFunFacts(t *testing.T) {
		type test struct {
			input string
			want  []string
		}
	
		tests := []test{
			{input: "watermelon", want: []string{"Watermelons are 92% water", "The first recorded watermelon harvest occurred nearly 5,000 years ago in Egypt"}},
			{input: "pumpkin", want: []string{"Pumpkins are a fruit, not a vegetable", "Pumpkins are 90% water", "The largest pumpkin ever grown weighed over 1,000 pounds"}},
		}
		
		for _, tc := range tests {
			got := GetFunFacts(tc.input)
			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("input: %v, expected: %v, got: %v", tc.input, tc.want, got)
			}
		}
	}
	

	for _, tc := range tests {
		got := GetFunFacts(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}|
}