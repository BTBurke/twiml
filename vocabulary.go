package twiml

import (
	"encoding/xml"
	"fmt"
)

// Twilio Client TwiML
type Client struct {
	XMLName  xml.Name `xml:"Client"`
	Name     string   `xml:",chardata"`
	Identity string   `xml:"Identity,omitempty"` // same as name

	Method               string   `xml:"method,attr,omitempty"`
	URL                  string   `xml:"url,attr,omitempty"`
	StatusCallback       string   `xml:"statusCallback,attr,omitempty"`
	StatusCallbackEvent  string   `xml:"statusCallbackEvent,attr,omitempty"`
	StatusCallbackMethod string   `xml:"statusCallbackMethod,attr,omitempty"`
	Children             []Markup `xml:",omitempty"`
}

// Add adds noun structs to a Client response as children
func (c *Client) Add(ml ...Markup) {
	for _, s := range ml {
		c.Children = append(c.Children, s)
	}
	return
}

// Validate returns an error if the TwiML is constructed improperly
func (c *Client) Validate() error {
	var ok bool

	if ok = Validate(AllowedMethod(c.Method)); !ok {
		return fmt.Errorf("Client markup failed validation")
	}

	// require either name or identity
	if ok = len(c.Name) > 0 || len(c.Identity) > 0; !ok {
		return fmt.Errorf("Client markup failed validation")
	}

	if ok = c.validParameters(); !ok {
		return fmt.Errorf("Client markup failed validation")
	}

	return nil
}

// validParameters checks that if parameters are set, name is empty and we have an identity
func (c *Client) validParameters() bool {
	// cannot have both of these be true at once
	return len(c.Children) == 0 || (len(c.Name) == 0 || len(c.Identity) != 0)
}

// Type returns the XML name of the verb
func (c *Client) Type() string {
	return "Client"
}

// Twilio Client Parameter TwiML
type Parameter struct {
	XMLName xml.Name `xml:"Parameter"`
	Name    string   `xml:"name,attr,omitempty"`
	Value   string   `xml:"value,attr,omitempty"`
}

func (p Parameter) Type() string {
	return "Parameter"
}

// Validate <Parameter> noun
func (p Parameter) Validate() error {
	if ok := Validate(Required(p.Name), Required(p.Value)); !ok {
		return fmt.Errorf("parameter markup failed validation")
	}

	return nil
}

// Conference TwiML
type Conference struct {
	XMLName                       xml.Name `xml:"Conference"`
	ConferenceName                string   `xml:",chardata"`
	Muted                         bool     `xml:"muted,attr,omitempty"`
	Beep                          string   `xml:"beep,attr,omitempty"`
	StartConferenceOnEnter        bool     `xml:"startConferenceOnEnter,attr,omitempty"`
	EndConferenceOnExit           bool     `xml:"endConferenceOnExit,attr,omitempty"`
	WaitURL                       string   `xml:"waitUrl,attr,omitempty"`
	WaitMethod                    string   `xml:"waitMethod,attr,omitempty"`
	MaxParticipants               int      `xml:"maxParticipants,attr,omitempty"`
	Record                        string   `xml:"record,attr,omitempty"`
	Region                        string   `xml:"region,attr,omitempty"`
	Trim                          string   `xml:"trim,attr,omitempty"`
	StatusCallbackEvent           string   `xml:"statusCallbackEvent,attr,omitempty"`
	StatusCallback                string   `xml:"statusCallback,attr,omitempty"`
	StatusCallbackMethod          string   `xml:"statusCallbackMethod,attr,omitempty"`
	RecordingStatusCallback       string   `xml:"recordingStatusCallback,attr,omitempty"`
	RecordingStatusCallbackMethod string   `xml:"recordingStatusCallbackMethod,attr,omitempty"`
	EventCallbackURL              string   `xml:"eventCallbackUrl,attr,omitempty"`
}

// Validate returns an error if the TwiML is constructed improperly
func (c *Conference) Validate() error {
	ok := Validate(
		OneOfOpt(c.Beep, "true", "false", "onEnter", "onExit"),
		AllowedMethod(c.WaitMethod),
		OneOfOpt(c.Record, "do-not-record", "record-from-start"),
		OneOfOpt(c.Trim, "trim-silence", "do-not-trim"),
		AllowedCallbackEvent(c.StatusCallbackEvent, ConferenceCallbackEvents),
		AllowedMethod(c.StatusCallbackMethod),
		AllowedMethod(c.RecordingStatusCallbackMethod),
	)
	if !ok {
		return fmt.Errorf("Conference markup failed validation")
	}
	return nil
}

// Type returns the XML name of the verb
func (c *Conference) Type() string {
	return "Conference"
}

// Dial TwiML
type Dial struct {
	XMLName xml.Name `xml:"Dial"`

	Action         string `xml:"action,attr,omitempty"`
	AnswerOnBridge bool   `xml:"answerOnBridge,attr,omitempty"`
	CallerID       string `xml:"callerId,attr,omitempty"`
	HangupOnStar   bool   `xml:"hangupOnStar,attr,omitempty"`
	Method         string `xml:"method,attr,omitempty"`

	Record                        string `xml:"record,attr,omitempty"`
	RecordingStatusCallback       string `xml:"recordingStatusCallback,attr,omitempty"`
	RecordingStatusCallbackMethod string `xml:"recordingStatusCallbackMethod,attr,omitempty"`
	RecordingStatusCallbackEvent  string `xml:"recordingStatusCallbackEvent,attr,omitempty"`

	RingTone string `xml:"ringTone,attr,omitempty"`

	Timeout   int `xml:"timeout,attr,omitempty"`
	TimeLimit int `xml:"timeLimit,attr,omitempty"`

	Trim string `xml:"trim,attr,omitempty"`

	Number   string   `xml:",chardata"`
	Children []Markup `xml:",omitempty"`
}

// Validate returns an error if the TwiML is constructed improperly
func (d *Dial) Validate() error {
	var errs []error
	for _, s := range d.Children {
		switch t := s.Type(); t {
		default:
			return fmt.Errorf("Not a valid verb under Dial: '%T'", s)
		case "Client", "Conference", "Number", "Queue", "Sip":
			if childErr := s.Validate(); childErr != nil {
				errs = append(errs, childErr)
			}
		}
	}

	ok := Validate(
		OneOfOpt(d.Method, "GET", "POST"),
	)
	if !ok {
		errs = append(errs, fmt.Errorf("Dial did not pass validation"))
	}

	if len(errs) > 0 {
		return ValidationError{errs}
	}
	return nil
}

// Add adds noun structs to a Dial response as children
func (d *Dial) Add(ml ...Markup) {
	for _, s := range ml {
		d.Children = append(d.Children, s)
	}
	return
}

// Type returns the XML name of the verb
func (d *Dial) Type() string {
	return "Dial"
}

// Enqueue TwiML
type Enqueue struct {
	XMLName       xml.Name `xml:"Enqueue"`
	Action        string   `xml:"action,attr,omitempty"`
	Method        string   `xml:"method,attr,omitempty"`
	WaitURL       string   `xml:"waitUrl,attr,omitempty"`
	WaitURLMethod string   `xml:"waitUrlMethod,attr,omitempty"`
	WorkflowSid   string   `xml:"workflowSid,attr,omitempty"`
	QueueName     string   `xml:",chardata"`
}

// Validate returns an error if the TwiML is constructed improperly
func (e *Enqueue) Validate() error {
	ok := Validate(
		AllowedMethod(e.Method),
		AllowedMethod(e.WaitURLMethod),
	)
	if !ok {
		return fmt.Errorf("%s markup failed validation", e.Type())
	}
	return nil
}

// Type returns the XML name of the verb
func (e *Enqueue) Type() string {
	return "Enqueue"
}

// Hangup TwiML
type Hangup struct {
	XMLName xml.Name `xml:"Hangup"`
}

// Validate returns an error if the TwiML is constructed improperly
func (h *Hangup) Validate() error {
	return nil
}

// Type returns the XML name of the verb
func (h *Hangup) Type() string {
	return "Hangup"
}

// Leave TwiML
type Leave struct {
	XMLName xml.Name `xml:"Leave"`
}

// Validate returns an error if the TwiML is constructed improperly
func (l *Leave) Validate() error {
	return nil
}

// Type returns the XML name of the verb
func (l *Leave) Type() string {
	return "Leave"
}

// Sms TwiML sends an SMS message. Text is required.  See the Twilio docs
// for an explanation of the default values of to and from.
type Sms struct {
	XMLName        xml.Name `xml:"Message"`
	To             string   `xml:"to,attr,omitempty"`
	From           string   `xml:"from,attr,omitempty"`
	Action         string   `xml:"action,attr,omitempty"`
	Method         string   `xml:"method,attr,omitempty"`
	StatusCallback string   `xml:"statusCallback,attr,omitempty"`
	Text           string   `xml:",chardata"`
}

// Validate returns an error if the TwiML is constructed improperly
func (s *Sms) Validate() error {
	ok := Validate(
		AllowedMethod(s.Method),
		Required(s.Text),
	)
	if !ok {
		return fmt.Errorf("%s markup failed validation", s.Type())
	}
	return nil
}

// Type returns the XML name of the verb
func (s *Sms) Type() string {
	return "Sms"
}

// Number TwiML
type Number struct {
	XMLName    xml.Name `xml:"Number"`
	SendDigits string   `xml:"sendDigits,attr,omitempty"`
	URL        string   `xml:"url,attr,omitempty"`
	Method     string   `xml:"method,attr,omitempty"`
	Number     string   `xml:",chardata"`
}

// Validate returns an error if the TwiML is constructed improperly
func (n *Number) Validate() error {
	ok := Validate(
		NumericOpt(n.SendDigits),
		AllowedMethod(n.Method),
		Required(n.Number),
	)
	if !ok {
		return fmt.Errorf("%s markup failed validation", n.Type())
	}
	return nil
}

// Type returns the XML name of the verb
func (n *Number) Type() string {
	return "Number"
}

// Pause TwiML
type Pause struct {
	XMLName xml.Name `xml:"Pause"`
	Length  int      `xml:"length,attr,omitempty"`
}

// Validate returns an error if the TwiML is constructed improperly
func (p *Pause) Validate() error {
	return nil
}

// Type returns the XML name of the verb
func (p *Pause) Type() string {
	return "Pause"
}

// Play TwiML
type Play struct {
	XMLName xml.Name `xml:"Play"`
	Loop    int      `xml:"loop,attr,omitempty"`
	Digits  string   `xml:"digits,attr,omitempty"`
	URL     string   `xml:",chardata"`
}

// Validate returns an error if the TwiML is constructed improperly
func (p *Play) Validate() error {

	ok := Validate(
		Required(p.URL),
		NumericOrWait(p.Digits),
	)

	if !ok {
		return fmt.Errorf("%s markup failed validation", p.Type())
	}
	return nil
}

// Type returns the XML name of the verb
func (p *Play) Type() string {
	return "Play"
}

// Queue TwiML
type Queue struct {
	XMLName             xml.Name `xml:"Queue"`
	URL                 string   `xml:"url,attr,omitempty"`
	Method              string   `xml:"method,attr,omitempty"`
	ReservationSid      string   `xml:"reservationSid,attr,omitempty"`
	PostWorkActivitySid string   `xml:"postWorkActivitySid,attr,omitempty"`
	Name                string   `xml:",chardata"`
}

// Validate returns an error if the TwiML is constructed improperly
func (q *Queue) Validate() error {
	ok := Validate(
		AllowedMethod(q.Method),
		Required(q.Name),
	)
	if !ok {
		return fmt.Errorf("%s markup failed validation", q.Type())
	}
	return nil
}

// Type returns the XML name of the verb
func (q *Queue) Type() string {
	return "Queue"
}

// Record TwiML
type Record struct {
	XMLName                       xml.Name `xml:"Record"`
	Action                        string   `xml:"action,attr,omitempty"`
	Method                        string   `xml:"method,attr,omitempty"`
	Timeout                       int      `xml:"timeout,attr,omitempty"`
	FinishOnKey                   string   `xml:"finishOnKey,attr,omitempty"`
	MaxLength                     int      `xml:"maxLength,attr,omitempty"`
	PlayBeep                      bool     `xml:"playBeep,attr,omitempty"`
	Trim                          string   `xml:"trim,attr,omitempty"`
	RecordingStatusCallback       string   `xml:"recordingStatusCallback,attr,omitempty"`
	RecordingStatusCallbackMethod string   `xml:"recordingStatusCallbackMethod,attr,omitempty"`
	Transcribe                    bool     `xml:"transcribe,attr,omitempty"`
	TranscribeCallback            string   `xml:"transcribeCallback,attr,omitempty"`
}

// Validate returns an error if the TwiML is constructed improperly
func (r *Record) Validate() error {
	ok := Validate(
		AllowedMethod(r.Method),
		OneOfOpt(r.Trim, TrimSilence, DoNotTrim),
		AllowedMethod(r.RecordingStatusCallbackMethod),
	)
	if !ok {
		return fmt.Errorf("%s markup failed validation", r.Type())
	}
	return nil
}

// Type returns the XML name of the verb
func (r *Record) Type() string {
	return "Record"
}

// Redirect TwiML
type Redirect struct {
	XMLName xml.Name `xml:"Redirect"`
	Method  string   `xml:"method,attr,omitempty"`
	URL     string   `xml:",chardata"`
}

// Validate returns an error if the TwiML is constructed improperly
func (r *Redirect) Validate() error {
	ok := Validate(
		AllowedMethod(r.Method),
		Required(r.URL),
	)
	if !ok {
		return fmt.Errorf("%s markup failed validation", r.Type())
	}
	return nil
}

// Type returns the XML name of the verb
func (r *Redirect) Type() string {
	return "Redirect"
}

// Reject TwiML
type Reject struct {
	XMLName xml.Name `xml:"Reject"`
	Reason  string   `xml:"reason,attr,omitempty"`
}

// Validate returns an error if the TwiML is constructed improperly
func (r *Reject) Validate() error {
	ok := Validate(
		OneOfOpt(r.Reason, "rejected", "busy"),
	)
	if !ok {
		return fmt.Errorf("%s markup failed validation", r.Type())
	}
	return nil
}

// Type returns the XML name of the verb
func (r *Reject) Type() string {
	return "Reject"
}

// Say TwiML
type Say struct {
	XMLName  xml.Name `xml:"Say"`
	Voice    string   `xml:"voice,attr,omitempty"`
	Language string   `xml:"language,attr,omitempty"`
	Loop     int      `xml:"loop,attr,omitempty"`
	Text     string   `xml:",chardata"`
}

// Validate returns an error if the TwiML is constructed improperly
func (s *Say) Validate() error {
	ok := Validate(
		OneOfOpt(s.Voice, Man, Woman, Alice),
		AllowedLanguage(s.Voice, s.Language),
		Required(s.Text),
	)
	if !ok {
		return fmt.Errorf("%s markup failed validation", s.Type())
	}
	return nil

}

// Type returns the XML name of the verb
func (s *Say) Type() string {
	return "Say"
}

// Sip TwiML
type Sip struct {
	XMLName              xml.Name `xml:"Sip"`
	Username             string   `xml:"username,attr,omitempty"`
	Password             string   `xml:"password,attr,omitempty"`
	URL                  string   `xml:"url,attr,omitempty"`
	Method               string   `xml:"method,attr,omitempty"`
	StatusCallbackEvent  string   `xml:"statusCallbackEvent,attr,omitempty"`
	StatusCallback       string   `xml:"statusCallback,attr,omitempty"`
	StatusCallbackMethod string   `xml:"statusCallbackMethod,attr,omitempty"`
	Address              string   `xml:",chardata"`
}

// TODO: Needs helpers to construct the SIP URL (specifying transport
// and headers) See https://www.twilio.com/docs/api/twiml/sip

// Validate returns an error if the TwiML is constructed improperly
func (s *Sip) Validate() error {
	//TODO: needs a custom validator type for statusCallbackEvent when set
	//because valid values can be concatenated
	ok := Validate(
		AllowedMethod(s.StatusCallbackMethod),
		AllowedCallbackEvent(s.StatusCallbackEvent, SipCallbackEvents),
		Required(s.Address),
	)
	if !ok {
		return fmt.Errorf("%s markup failed validation", s.Type())
	}
	return nil
}

// Type returns the XML name of the verb
func (s *Sip) Type() string {
	return "Sip"
}

// Gather TwiML
type Gather struct {
	XMLName               xml.Name `xml:"Gather"`
	Action                string   `xml:"action,attr,omitempty"`
	Method                string   `xml:"method,attr,omitempty"`
	Timeout               int      `xml:"timeout,attr,omitempty"`
	FinishOnKey           string   `xml:"finishOnKey,attr,omitempty"`
	NumDigits             int      `xml:"numDigits,attr,omitempty"`
	Input                 string   `xml:"input,attr,omitempty"`
	Hints                 string   `xml:"hints,attr,omitempty"`
	PartialResultCallback string   `xml:"partialResultCallback,attr,omitempty"`
	Language              string   `xml:"language,attr,omitempty"`
	ProfanityFilter       bool     `xml:"profanityFilter,attr,omitempty"`
	SpeechTimeout         int      `xml:"speechTimeout,attr,omitempty"`
	Children              []Markup `valid:"-"`
}

// Validate returns an error if the TwiML is constructed improperly
func (g *Gather) Validate() error {
	var errs []error

	for _, s := range g.Children {
		switch t := s.Type(); t {
		default:
			return fmt.Errorf("Not a valid verb as child of Gather: '%T'", s)
		case "Say", "Play", "Pause":
			if childErr := s.Validate(); childErr != nil {
				errs = append(errs, childErr)
			}
		}
	}
	ok := Validate(
		AllowedMethod(g.Method),
	)
	if !ok {
		errs = append(errs, fmt.Errorf("Gather failed validation"))
		return ValidationError{errs}
	}
	if len(errs) > 0 {
		return ValidationError{errs}
	}
	return nil
}

// Add collects digits a caller enter by pressing the keypad to an existing Gather verb.
// Valid nested verbs: Say, Pause, Play
func (g *Gather) Add(ml ...Markup) {
	for _, s := range ml {
		g.Children = append(g.Children, s)
	}
	return
}

// Type returns the XML name of the verb
func (g *Gather) Type() string {
	return "Gather"
}
