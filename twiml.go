// Package twiml provides Twilio Markup Language support for building web
// services with instructions for twilio how to handle incoming call or message.
package twiml

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
)

// Response container for other TwiML verbs
type Response struct {
	XMLName  xml.Name `xml:"Response"`
	Response []interface{}
}

// NewResponse creates new response
func NewResponse() *Response {
	return new(Response)
}

// Add appends TwiML verb structs to response. Valid verbs: Enqueue, Say,
// Leave, Message, Pause, Play, Record, Redirect, Reject, Hangup
func (r *Response) Add(structs ...interface{}) error {
	for _, s := range structs {
		switch s := s.(type) {
		default:
			return fmt.Errorf("Not a valid verb to return in response: '%T'", s)
		case Enqueue, Hangup, Leave, Pause, Play, Record,
			Redirect, Reject, Say, Dial, Gather:
			r.Response = append(r.Response, s)
		}
	}
	return nil
}

// Write sends XML encoded response to writer
func (r Response) Write(w io.Writer) error {
	if len(r.Response) == 0 {
		return fmt.Errorf("Can not encode an empty response")
	}

	enc := xml.NewEncoder(w)
	enc.Indent("", "  ")

	_, err := w.Write([]byte(xml.Header))
	if err != nil {
		return err
	}

	if err := enc.Encode(r); err != nil {
		return err
	}
	return nil
}

// String returns a formatted XML response
func (r Response) String() (string, error) {
	var buf = new(bytes.Buffer)
	if err := r.Write(buf); err != nil {
		return "", err
	}
	return buf.String(), nil
}
