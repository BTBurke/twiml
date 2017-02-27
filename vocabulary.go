package twiml

import (
	"encoding/xml"
	"fmt"
)

// Client TwiML
type Client struct {
	XMLName xml.Name `xml:"Client"`
	Method  string   `xml:"method,attr,omitempty"`
	URL     string   `xml:"URL,omitempty"`
	Name    string   `xml:",chardata"`
}

// Conference TwiML
type Conference struct {
	XMLName                xml.Name `xml:"Conference"`
	Muted                  bool     `xml:"muted,attr,omitempty"`
	Beep                   string   `xml:"beep,attr,omitempty"`
	StartConferenceOnEnter bool     `xml:"startConferenceOnEnter,attr,omitempty"`
	EndConferenceOnExit    bool     `xml:"endConferenceOnExit,attr,omitempty"`
	WaitURL                string   `xml:"waitUrl,attr,omitempty"`
	WaitMethod             string   `xml:"waitMethod,attr,omitempty"`
	MaxParticipants        int      `xml:"maxParticipants,attr,omitempty"`
	Name                   string   `xml:",chardata"`
}

// Dial TwiML
type Dial struct {
	XMLName      xml.Name      `xml:"Dial"`
	Action       string        `xml:"action,attr,omitempty"`
	Method       string        `xml:"method,attr,omitempty"`
	Timeout      int           `xml:"timeout,attr,omitempty"`
	HangupOnStar bool          `xml:"hangupOnStar,attr,omitempty"`
	TimeLimit    int           `xml:"timeLimit,attr,omitempty"`
	CallerID     string        `xml:"callerId,attr,omitempty"`
	Record       bool          `xml:"record,attr,omitempty"`
	Number       string        `xml:",chardata"`
	Nested       []interface{} `xml:",omitempty"`
}

// Add appends noun structs to a Dial response
// Valid nouns: Client, Conference, Number, Queue, Sip
func (d *Dial) Add(structs ...interface{}) error {

	for _, s := range structs {
		switch s := s.(type) {
		default:
			return fmt.Errorf("Not a valid verb under Dial: '%T'", s)
		case Client, Conference, Number, Queue, Sip:
			d.Nested = append(d.Nested, s)
		}
	}
	return nil
}

// Enqueue TwiML
type Enqueue struct {
	XMLName       xml.Name `xml:"Enqueue"`
	Action        string   `xml:"action,attr,omitempty"`
	Method        string   `xml:"method,attr,omitempty"`
	WaitURL       string   `xml:"waitUrl,attr,omitempty"`
	WaitURLMethod string   `xml:"waiUrlMethod,attr,omitempty"`
	Name          string   `xml:",chardata"`
}

// Hangup TwiML
type Hangup struct {
	XMLName xml.Name `xml:"Hangup"`
}

// Leave TwiML
type Leave struct {
	XMLName xml.Name `xml:"Leave"`
}

// Message TwiML
type Message struct {
	XMLName        xml.Name `xml:"Message"`
	To             string   `xml:"to,attr,omitempty"`
	From           string   `xml:"from,attr,omitempty"`
	Action         string   `xml:"action,attr,omitempty"`
	Method         string   `xml:"method,attr,omitempty"`
	StatusCallback string   `xml:"statusCallback,attr,omitempty"`
	Body           string   `xml:"Body,omitempty"`
	Media          string   `xml:"Media,omitempty"`
}

// Number TwiML
type Number struct {
	XMLName    xml.Name `xml:"Number"`
	SendDigits string   `xml:"sendDigits,attr,omitempty"`
	URL        string   `xml:"url,attr,omitempty"`
	Method     string   `xml:"method,attr,omitempty"`
	Number     string   `xml:",chardata"`
}

// Pause TwiML
type Pause struct {
	XMLName xml.Name `xml:"Pause"`
	Length  int      `xml:"length,attr,omitempty"`
}

// Play TwiML
type Play struct {
	XMLName xml.Name `xml:"Play"`
	Loop    int      `xml:"loop,attr,omitempty"`
	Digits  int      `xml:"digits,attr,omitempty"`
	URL     string   `xml:",chardata"`
}

// Queue TwiML
type Queue struct {
	XMLName xml.Name `xml:"Queue"`
	URL     string   `xml:"url,attr,omitempty"`
	Method  string   `xml:"method,attr,omitempty"`
	Name    string   `xml:",chardata"`
}

// Record TwiML
type Record struct {
	XMLName            xml.Name `xml:"Record"`
	Action             string   `xml:"action,attr,omitempty"`
	Method             string   `xml:"method,attr,omitempty"`
	Timeout            int      `xml:"timeout,attr,omitempty"`
	FinishOnKey        string   `xml:"finishOnKey,attr,omitempty"`
	MaxLength          int      `xml:"maxLength,attr,omitempty"`
	Transcribe         bool     `xml:"transcribe,attr,omitempty"`
	TranscribeCallback string   `xml:"transcribeCallback,attr,omitempty"`
	PlayBeep           bool     `xml:"playBeep,attr,omitempty"`
}

// Redirect TwiML
type Redirect struct {
	XMLName xml.Name `xml:"Redirect"`
	Method  string   `xml:"method,attr,omitempty"`
	URL     string   `xml:",chardata"`
}

// Reject TwiML
type Reject struct {
	XMLName xml.Name `xml:"Reject"`
	Reason  string   `xml:"reason,attr,omitempty"`
}

// Say TwiML
type Say struct {
	XMLName  xml.Name `xml:"Say"`
	Voice    string   `xml:"voice,attr,omitempty"`
	Language string   `xml:"language,attr,omitempty"`
	Loop     int      `xml:"loop,attr,omitempty"`
	Text     string   `xml:",chardata"`
}

// Sip TwiML
type Sip struct {
	XMLName  xml.Name `xml:"Sip"`
	Username string   `xml:"username,attr,omitempty"`
	Password string   `xml:"password,attr,omitempty"`
	URL      string   `xml:"url,attr,omitempty"`
	Method   string   `xml:"method,attr,omitempty"`
	Address  string   `xml:",chardata"`
}

// Gather TwiML
type Gather struct {
	XMLName     xml.Name `xml:"Gather"`
	Action      string   `xml:"action,attr,omitempty"`
	Method      string   `xml:"method,attr,omitempty"`
	Timeout     int      `xml:"timeout,attr,omitempty"`
	FinishOnKey string   `xml:"finishOnKey,attr,omitempty"`
	NumDigits   int      `xml:"numDigits,attr,omitempty"`
	Nested      []interface{}
}

// Add collects digits a caller enter by pressing the keypad to an existing Gather verb.
// Valid nested verbs: Say, Pause, Play
func (g *Gather) Add(structs ...interface{}) error {
	for _, s := range structs {
		switch s := s.(type) {
		default:
			return fmt.Errorf("Not a valid verb under Gather: '%T'", s)
		case Say, Pause, Play: // Valid nested verbs
			g.Nested = append(g.Nested, s)
		}

	}
	return nil
}
