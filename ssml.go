package twiml

import (
	"encoding/xml"
	"fmt"
)

// SSML Options
const (
	// https://developer.amazon.com/docs/custom-skills/speech-synthesis-markup-language-ssml-reference.html#break
	SSMLNone    = "none"
	SSMLXWeak   = "x-weak"
	SSMLWeak    = "weak"
	SSMLMedium  = "medium"
	SSMLStrong  = "strong"
	SSMLXStrong = "x-strong"

	// https://developer.amazon.com/docs/custom-skills/speech-synthesis-markup-language-ssml-reference.html#say-as
	SSMLSpellOut     = "spell-out"
	SSMLCharacters   = "characters"
	SSMLOrdinal      = "ordinal"
	SSMLDigits       = "digits"
	SSMLFraction     = "fraction"
	SSMLUnit         = "unit"
	SSMLDate         = "date"
	SSMLTime         = "time"
	SSMLTelephone    = "telephone"
	SSMLAddress      = "address"
	SSMLInterjection = "interjection"
	SSMLExpletive    = "expletive"

	// https://developer.amazon.com/docs/custom-skills/speech-synthesis-markup-language-ssml-reference.html#emphasis
	SSMLModerate = "moderate"
	SSMLReduced  = "reduced"

	// https://developer.amazon.com/docs/custom-skills/speech-synthesis-markup-language-ssml-reference.html#prosody
	SSMLXSlow = "x-slow"
	SSMLSlow  = "slow"
	SSMLFast  = "fast"
	SSMLXFast = "x-fast"

	SSMLXLow  = "x-low"
	SSMLLow   = "low"
	SSMLHigh  = "high"
	SSMLXHigh = "x-high"

	SSMLSilent = "silent"
	SSMLXSoft  = "x-soft"
	SSMLSoft   = "soft"
	SSMLLoud   = "loud"
	SSMLXLoud  = "x-loud"

	// https://developer.amazon.com/docs/custom-skills/speech-synthesis-markup-language-ssml-reference.html#amazon-effect
	SSMLEffectWhisper = "whispered"
)

// SSMLText
type SSMLText struct {
	XMLName  xml.Name `xml:"p"`
	Text    string   `xml:",chardata"`
}

// Validate returns an error if the TwiML is constructed improperly
func (c *SSMLText) Validate() error {
	return nil
}

// Type returns the XML name of the verb
func (c *SSMLText) Type() string {
	return "SSMLText"
}


// SSMLBreak
type SSMLBreak struct {
	XMLName  xml.Name `xml:"break"`
	Strength string   `xml:"strength,attr,omitempty"`
	Time     string   `xml:"time,attr,omitempty"`
}

// Validate returns an error if the TwiML is constructed improperly
func (c *SSMLBreak) Validate() error {
	ok := Validate(
		OneOfOpt(c.Strength, SSMLNone, SSMLXWeak, SSMLWeak, SSMLMedium, SSMLStrong, SSMLXStrong),
	)
	if !ok {
		return fmt.Errorf("%s markup failed validation", c.Type())
	}
	return nil
}

// Type returns the XML name of the verb
func (c *SSMLBreak) Type() string {
	return "SSMLBreak"
}

// SSMLSayAs
type SSMLSayAs struct {
	XMLName     xml.Name `xml:"say-as"`
	InterpretAs string   `xml:"interpret-as,attr,omitempty"`
	Text        string   `xml:",chardata"`
}

// Validate returns an error if the TwiML is constructed improperly
func (c *SSMLSayAs) Validate() error {
	ok := Validate(
		OneOfOpt(c.InterpretAs, SSMLSpellOut, SSMLCharacters, SSMLOrdinal, SSMLDigits, SSMLFraction, SSMLUnit, SSMLDate, SSMLTime, SSMLTelephone, SSMLAddress, SSMLInterjection, SSMLExpletive),
		Required(c.Text),
	)
	if !ok {
		return fmt.Errorf("%s markup failed validation", c.Type())
	}
	return nil
}

// Type returns the XML name of the verb
func (c *SSMLSayAs) Type() string {
	return "SSMLSayAs"
}

// SSMLEmphasis
type SSMLEmphasis struct {
	XMLName xml.Name `xml:"emphasis"`
	Level   string   `xml:"level,attr,omitempty"`
	Text    string   `xml:",chardata"`
}

// Validate returns an error if the TwiML is constructed improperly
func (c *SSMLEmphasis) Validate() error {
	ok := Validate(
		OneOfOpt(c.Level, SSMLStrong, SSMLModerate, SSMLReduced),
		Required(c.Text),
	)
	if !ok {
		return fmt.Errorf("%s markup failed validation", c.Type())
	}
	return nil
}

// Type returns the XML name of the verb
func (c *SSMLEmphasis) Type() string {
	return "SSMLEmphasis"
}

// SSMLProsody
type SSMLProsody struct {
	XMLName xml.Name `xml:"prosody"`
	Rate    string   `xml:"rate,attr,omitempty"`
	Volume  string   `xml:"volume,attr,omitempty"`
	Pitch   string   `xml:"pitch,attr,omitempty"`
	Text    string   `xml:",chardata"`
}

// Validate returns an error if the TwiML is constructed improperly
func (c *SSMLProsody) Validate() error {
	ok := Validate(
		Required(c.Text),
	)
	if !ok {
		return fmt.Errorf("%s markup failed validation", c.Type())
	}
	return nil
}

// Type returns the XML name of the verb
func (c *SSMLProsody) Type() string {
	return "SSMLProsody"
}

// SSMLEffect
// https://developer.amazon.com/docs/custom-skills/speech-synthesis-markup-language-ssml-reference.html#amazon-effect
// Amazon Specific SSML - must use Polly voices
type SSMLEffect struct {
	XMLName xml.Name `xml:"amazon:effect"`
	Name    string   `xml:"name,attr,omitempty"`
	Text    string   `xml:",chardata"`
}

// Validate returns an error if the TwiML is constructed improperly
func (c *SSMLEffect) Validate() error {
	ok := Validate(
		OneOfOpt(c.Name, SSMLEffectWhisper),
		Required(c.Text),
	)
	if !ok {
		return fmt.Errorf("%s markup failed validation", c.Type())
	}
	return nil
}

// Type returns the XML name of the verb
func (c *SSMLEffect) Type() string {
	return "SSMLEffect"
}
