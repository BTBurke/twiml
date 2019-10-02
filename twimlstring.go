package twiml

import (
	"regexp"
	"strconv"
	"strings"
)

// TwiMLStrings is a simple string macro language allowing to define TwiML like verbs directly in a string.
// It follows a simple format where a TwiML statement is encapsulated within brackets and the verb divided
// by an | from it's optional content and attributes. The content can be encapsulated with single backward
// quotes (`) to ensure easy compatibility with JSON. Attributes are separated by comma and follow the
// format key:value.
//
// Format:
// {<verb>|<content>,<attr:key>:<attr:value>,...}
// {<verb>|`<content>`,<attr:key>:<attr:value>,...}
//
// Examples:
// ts := "{say|Hello World}"
// ts := "{say|Hello}{strong|World}{say|Hello World,voice:man}
// ts := "{say|`Hello, World`,voice:female}
//
// It tries to mimic as close as possible the TWIML verbs with their associated attributes however implements
// assorted new verbs that allow quick use of advance SSML techniques such as saying an address or telephone number.
//
// ts := "{say|Call me back at}{telephone|(949)555-1234}"
//
//

// ParseString ...
func ParseString(s string) []Markup   {

	var out []Markup

	re := regexp.MustCompile("\\{(.*?)\\}")
	matches := re.FindAllStringSubmatch(s, -1)

	if matches != nil {

		for _, m := range matches {

			statement := strings.Split(m[1], "|")
			if len(statement) > 0 {

				pp := []string{}
				if len(statement) > 1 {

					attr := statement[1];
					ru := regexp.MustCompile("\\`(.*?)\\`")
					rm := ru.FindStringSubmatch(s)

					if rm != nil {
						attr = strings.Replace(attr, rm[1] + ",", "", -1)
					}
					pp = strings.Split(attr, ",")
					if rm != nil {
						pp = append([]string{rm[1]}, pp...)
					}
				}

				p := toTwiml(statement[0], pp)

				if p != nil {
					out = append(out, p)
				}
			}
		}

	} else {
		out = append(out, &Say{Text: s})
	}

	return out
}

// toTwiml ...
func toTwiml(statement string, attributes []string) Markup {

	content := "";
	if len(attributes) > 0 {
		content = attributes[0]
	}

	data := make(map[string]string)
	if len(attributes) > 1 {
		params := attributes[1:]
		if len(params) > 0 {
			for _, p := range params {
				pp := strings.Split(p, ":")
				if len(pp) == 2 {
					data[pp[0]] = pp[1]
				}
			}
		}
	}



	switch statement {
	default:
		return nil
	case "p":
		return &Say{
			Children: []Markup{
				&SSMLText{
					Text: content,
				},
			},
		}
	case "whisper":
		return &Say{
			Children: []Markup{
				&SSMLEffect{
					Name: withDefault(data["effect"], SSMLEffectWhisper),
					Text: content,
				},
			},
		}
	case "strong":
		return &Say{
			Children: []Markup{
				&SSMLEmphasis{
					Text: content,
					Level: withDefault(data["level"], SSMLStrong),
				},
			},
		}
	case "pause":
		return &Say{
			Children: []Markup{
				&SSMLBreak{
					Time: content,
				},
			},
		}
	case SSMLSpellOut, SSMLCharacters, SSMLOrdinal, SSMLDigits, SSMLFraction, SSMLUnit, SSMLDate, SSMLTime, SSMLTelephone, SSMLAddress, SSMLInterjection, SSMLExpletive:
		return &Say{
			Children: []Markup{
				&SSMLSayAs{
					Text: content,
					InterpretAs: statement,
				},
			},
		}
	case "hangup":
		return &Hangup{}
	case "dtfm":
		return &Play{
			Digits: content,
		}
	case "play":
		return &Play{
			URL: content,
			Loop: asInt(data["loop"], 0),
		}
	case "say":
		return &Say{
			Text: content,
			Language: data["language"],
			Voice: data["voice"],
			Loop: asInt(data["loop"], 0),
			Polly: true,
		}
	case "dial":
		return &Dial{
			Number: content,
			Record: asInt(data["record"], 0) == 1,
		}
	}
}

func asInt(s string, d int) int {
	r, e := strconv.Atoi(s)
	if e != nil {
		return d
	}
	return r
}

func withDefault(s string, d string) string {
	if s == "" {
		return d
	}
	return s
}
