package twiml

import "regexp"

// Validate aggregates the results of individual validation functions and returns true
// when all validation functions pass
func Validate(vf ...bool) bool {
	for _, f := range vf {
		if !f {
			return false
		}
	}
	return true
}

// OneOf validates that a field is one of the options provided
func OneOf(field string, options ...string) bool {
	for _, w := range options {
		if field == w {
			return true
		}
	}
	return false
}

// IntBetween validates that a field is an integer between high and low
func IntBetween(field int, high int, low int) bool {
	if (field <= high) && (field >= low) {
		return true
	}
	return false
}

// Required validates that a field is not the empty string
func Required(field string) bool {
	if len(field) > 0 {
		return true
	}
	return false
}

// OneOfOpt validates that a field is one of the options provided or the empty string (for optional fields)
func OneOfOpt(field string, options ...string) bool {
	if field == "" {
		return true
	}
	return OneOf(field, options...)
}

// AllowedMethod validates that a method is either of type GET or POST (or empty string to default)
func AllowedMethod(field string) bool {
	// optional field always set with default (typically POST)
	if field == "" {
		return true
	}
	if (field != "GET") && (field != "POST") {
		return false
	}
	return true
}

// Numeric validates that a string contains only digits 0-9
func Numeric(field string) bool {
	matched, err := regexp.MatchString("^[0-9]+$", field)
	if err != nil {
		return false
	}
	return matched
}

// NumericOrWait validates that a string contains only digits 0-9 or the wait key 'w'
func NumericOrWait(field string) bool {
	matched, err := regexp.MatchString("^[0-9w]+$", field)
	if err != nil {
		return false
	}
	return matched
}

// NumericOpt validates that the field is numeric or empty string (for optional fields)
func NumericOpt(field string) bool {
	if field == "" {
		return true
	}
	return Numeric(field)
}

// AllowedLanguage validates that the combination of speaker and language is allowable
func AllowedLanguage(speaker string, language string) bool {
	switch speaker {
	case Man, Woman:
		return OneOfOpt(language, English, French, German, Spanish, EnglishUK)
	case Alice:
		return OneOfOpt(language,
			DanishDenmark,
			GermanGermany,
			EnglishAustralia,
			EnglishCanada,
			EnglishUK,
			EnglishIndia,
			EnglishUSA,
			SpanishCatalan,
			SpanishSpain,
			SpanishMexico,
			FinishFinland,
			FrenchCanada,
			FrenchFrance,
			ItalianItaly,
			JapaneseJapan,
			KoreanKorea,
			NorwegianNorway,
			DutchNetherlands,
			PolishPoland,
			PortugueseBrazil,
			PortuguesePortugal,
			RussianRussia,
			SwedishSweden,
			ChineseMandarin,
			ChineseCantonese,
			ChineseTaiwanese,
		)
	default:
		return OneOfOpt(language, English, French, German, Spanish, EnglishUK)
	}
}

// AllowedCallbackEvent validates that the CallbackEvent is one of the allowed options
func AllowedCallbackEvent(events string) bool {
	if events == "" {
		return true
	}
	var validEvents = regexp.MustCompile(`^(initiated\s?|ringing\s?|answered\s?|completed\s?)+$`)
	return validEvents.MatchString(events)
}
