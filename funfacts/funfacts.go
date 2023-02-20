package funfucts

/**
  Implementer funfacts-funksjon:
    GetFunFacts(about string) []string
      hvor about kan ha en av tre testverdier, -
        sun, luna eller terra
  Sett inn alle Funfucts i en struktur
  type FunFacts struct {
      Sun []string
      Luna []string
      Terra []string
  }
*/
type FunFacts struct {
	Sun   []string
	Luna  []string
	Terra []string
}

func GetFunFacts(about string) []string {
	facts := FunFacts{
		Sun: []string{
			"The Sun accounts for 99.86% of the mass in the Solar System.",
			"The Sun is about 109 times wider than the Earth.",
			"The temperature at the Sun's core is about 15 million degrees Celsius.",
		},
		Luna: []string{
			"The Moon is Earth's only natural satellite.",
			"The Moon is about 4.5 billion years old.",
			"The Moon is slowly moving away from Earth.",
		},
		Terra: []string{
			"Earth is the only planet known to support life.",
			"Earth is not a perfect sphere, it is slightly flattened at the poles.",
			"Earth has a powerful magnetic field that protects it from the solar wind.",
		},
	}

	switch about {
	case "sun":
		return facts.Sun
	case "luna":
		return facts.Luna
	case "terra":
		return facts.Terra
	default:
		return []string{}
	}
}
