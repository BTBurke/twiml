TwiML
===
A library for producing TwiML XML markup for use with the Twilio API.  This library can generate TwiML responses and provides helpers for processing callbacks and requests from Twilio.

This library does not yet cover the entire TwiML API, but pull requests are welcome if you find something missing.

## Processing a request from Twilio

The library contains helpers to bind incoming Twilio requests to a struct that includes all of the available info from the request.  Most initial requests from Twilio are of type `twiml.VoiceRequest`.  Other request types are possible as a result of callbacks you register in your response.  See the [GoDoc](https://godoc.org/BTBurke/twiml) for details.

```go
func(w http.ResponseWriter, r *http.Request) {
    var vr twiml.VoiceRequest
    if err := twiml.Bind(&vr, r); err != nil {
        http.Error(w, http.StatusText(400), 400)
        return
    }
    fmt.Printf("Incoming call from %s", vr.From)
}
```

## Constructing a response using TwiML

Once you receive a request from the Twilio API, you construct a TwiML response to provide directions for how to deal with the call.  This library includes (most of) the allowable verbs and rules to validate that your response is constructed properly.

``` 
// CallRequest will return XML to connect to the forwarding number
func CallRequest(cfg Config) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		
        // Bind the request
        var cr twiml.VoiceRequest
		if err := twiml.Bind(&cr, r); err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}

        // Create a new response container
		res := twiml.NewResponse()

		switch status := cr.CallStatus; status {
		
        // Call is already in progress, tell Twilio to continue
        case twiml.InProgress:
			w.WriteHeader(200)
			return
        
        // Call is ringing but has not been connected yet, respond with
        // a forwarding number
		case twiml.Ringing, twiml.Queued:
			// Create a new Dial verb
            d := twiml.Dial{
				Number:   cfg.ForwardingNumber,
				Action:   "action/",
				Timeout:  15,
				CallerID: cr.To,
			}

            // Add the verb to the response
			res.Add(&d)
			
            // Validate and encode the response.  Validation is done
            // automatically before the response is encoded.
            b, err := res.Encode()
			if err != nil {
				http.Error(w, http.StatusText(502), 502)
				return
			}

            // Write the XML response to the http.ReponseWriter
			if _, err := w.Write(b); err != nil {
				http.Error(w, http.StatusText(502), 502)
				return
			}
			w.Header().Set("Content-Type", "application/xml")
			w.WriteHeader(200)
			return

        // Call is over, hang up
		default:
			res.Add(&twiml.Hangup{})
			b, err := res.Encode()
			if err != nil {
				http.Error(w, http.StatusText(502), 502)
				return
			}
			if _, err := w.Write(b); err != nil {
				http.Error(w, http.StatusText(502), 502)
				return
			}
			w.Header().Set("Content-Type", "application/xml")
			w.WriteHeader(200)
			return
		}
	}
}
```

The example above shows the general flow of constructing a response.  Start with creating a new response container, then use the `Add()` method to add a TwiML verb with its appropriate configuration.  Verbs that allow other verbs to be nested within them expose their own `Add()` method.  On the call to `Encode()` the complete response is validated to ensure that the response is properly configured.

## More examples

For a more detailed example of constructing a small TwiML response server, see my [Twilio Voice project](https://github.com/BTBurke/twilio-voice) which is a Google-voice clone that forwards calls to your number and handles transcribing voicemails.