package twiml

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ClientWithIdentity(t *testing.T) {
	response := NewResponse()
	dial := &Dial{Number: " "}
	response.Add(dial)

	client := &Client{
		Identity:             Alice,
		URL:                  "http://google.com/url",
		Method:               "POST",
		StatusCallback:       "initiated",
		StatusCallbackEvent:  "GET",
		StatusCallbackMethod: "http://google.com/status",
	}
	dial.Add(client)

	client.Add(Parameter{Name: "FirstName", Value: "Alice"})
	client.Add(Parameter{Name: "LastName", Value: "Smith"})

	b, err := response.Encode()
	assert.NoError(t, err)
	assert.NotEmpty(t, b)

	data := string(b)
	assert.Contains(t, data, `<Client method="POST" url="http://google.com/url" statusCallback="initiated" statusCallbackEvent="GET" statusCallbackMethod="http://google.com/status">`)
	assert.Contains(t, data, `<Identity>alice</Identity>`)
	assert.Contains(t, data, `<Parameter name="FirstName" value="Alice"></Parameter>`)
	assert.Contains(t, data, `<Parameter name="LastName" value="Smith"></Parameter>`)

	t.Log(string(b))
}

func Test_ClientWithoutIdentity(t *testing.T) {
	response := NewResponse()
	dial := &Dial{Number: " "}
	response.Add(dial)

	client := &Client{Name: Alice}
	dial.Add(client)

	b, err := response.Encode()
	assert.NoError(t, err)
	assert.NotEmpty(t, b)

	data := string(b)
	assert.Contains(t, data, `<Client>alice</Client>`)
}

func TestClient_Validate(t *testing.T) {
	type fields struct {
		Name     string
		Identity string
		Method   string
		Children []Markup
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{name: "Identity_No_Parameters", fields: fields{Identity: "Alice"}, wantErr: false},
		{name: "Identity_Parameters", fields: fields{Identity: "Alice", Children: []Markup{&Parameter{Name: "FirstName", Value: "Alice"}, &Parameter{Name: "LastName", Value: "Smith"}}}, wantErr: false},
		{name: "Name_No_Parameters", fields: fields{Name: "Alice"}, wantErr: false},
		{name: "Name_Parameters", fields: fields{Name: "Alice", Children: []Markup{&Parameter{Name: "FirstName", Value: "Alice"}, &Parameter{Name: "LastName", Value: "Smith"}}}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				Name:     tt.fields.Name,
				Identity: tt.fields.Identity,
				Method:   tt.fields.Method,
				Children: tt.fields.Children,
			}
			if err := c.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
