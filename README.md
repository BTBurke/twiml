TwiML
===
[![Go Doc](https://godoc.org/github.com/BTBurke/twiml?status.svg)](https://godoc.org/github.com/BTBurke/twiml)     [![CircleCI](https://circleci.com/gh/BTBurke/twiml.svg?style=svg)](https://circleci.com/gh/BTBurke/twiml)


A library for producing TwiML XML markup for use with the Twilio API. This library can generate TwiML responses and provides helpers for processing callbacks and requests from Twilio. 

This library also supports SSML and TwiMLStrings.

This library does not yet cover the entire TwiML API, but pull requests are welcome if you find something missing.

## Example

```golang
func SaySomething() string {
    res := twiml.NewResponse()
    
    res.Add(
        &twiml.Say{
            Text: "Hello World",
        }
    )
    
    b, err := res.Encode()
    if err != nil {
        return ""
    }
    
    return string(b)
}
```


## Supported Verbs

#### [`<Conference>`](https://www.twilio.com/docs/voice/twiml/conference)

All Conference attributes are supported.

```golang
m := twiml.Conference{
    ConferenceName: "MyName",
}
```


#### [`<Dial>`](https://www.twilio.com/docs/voice/twiml/dial)

All Dial attributes are supported.

```golang
m := twiml.Dial{
    Number: "+15551239876",
}
```


#### [`<Enqueue>`](https://www.twilio.com/docs/voice/twiml/enqueue)

All Enqueue attributes are supported.

```golang
m := twiml.Enqueue{
    Queue: "holding-pattern",
}
```


#### [`<Sms>`](https://www.twilio.com/docs/voice/twiml/sms)

All SMS attributes are supported.

```golang
m := twiml.Sms{
    Message: "Hello World",
    From: "+15551230987",
    To: "+14321230987,
}
```


#### [`<Number>`](https://www.twilio.com/docs/voice/twiml/number)

All Number attributes are supported.

```golang
m := twiml.Number{
    Number: "+14321230987,
}
```

#### [`<Pause>`](https://www.twilio.com/docs/voice/twiml/pause)

All Pause attributes are supported. You can also use `SSMLBreak` verb to create more fine tuned pauses in your speech output.

```golang
m := twiml.Pause{
    Length: 3,
}
```

#### [`<Play>`](https://www.twilio.com/docs/voice/twiml/play)

All Play attributes are supported.

```golang
m := twiml.Play{
    URL: "http://url-to-file/audio.mp3",
}
```

#### [`<Queue>`](https://www.twilio.com/docs/voice/twiml/queue)

All Queue attributes are supported.

```golang
m := twiml.Queue{
    Name: "my-queue",
}
```

#### [`<Record>`](https://www.twilio.com/docs/voice/twiml/record)

All Record attributes are supported.

```golang
m := twiml.Record{
}
```

#### [`<Redirect>`](https://www.twilio.com/docs/voice/twiml/redirect)

All Redirect attributes are supported.

```golang
m := twiml.Redirect{
    URL: "http://twiml-bucket/twiml.json",
}
```

#### [`<Reject>`](https://www.twilio.com/docs/voice/twiml/reject)

All Redirect attributes are supported.

```golang
m := twiml.Reject{
    Reason: "busy",
}
```

#### [`<Say>`](https://www.twilio.com/docs/voice/twiml/say)

All Say attributes are supported. Also SSML is supported when using the `Children` attribute.


```golang
m := twiml.Say{
    Text: "Hello World",
}
```

In order to use any of the Polly voices, you need to configure your Twilio account as well explicit enable it in the statement:


```golang
m := twiml.Say{
    Text: "Hello World",
    Voice: "Brian",
    Polly : true,
}
```


#### [`<Sip>`](https://www.twilio.com/docs/voice/twiml/sip)

All Sip attributes are supported.

```golang
m := twiml.Sip{
    Address: "my-sip.net",
}
```

#### [`<Gather>`](https://www.twilio.com/docs/voice/twiml/gather)

All Gather attributes are supported. Also supports nesting of `Say` and other verbs within itself.

```golang
m := twiml.Gather{
}
```


#### [`<Hangup>`](https://www.twilio.com/docs/voice/twiml/hangup)

Hangs up the call.

```golang
m := twiml.Hangup{}
```

#### [`<Hangup>`](https://www.twilio.com/docs/voice/twiml/leave)

Leaves a conference call.

```golang
m := twiml.Leave{}
```



## SSML

The library contains a set of meaningful SSML implementations as outlined below. SSML verbs are only allowed when nested within a `Say` verb.

### Example

```golang

res := twiml.NewResponse()
res.Add(&twiml.Say{
    Children: []twiml.Markup{
        &twiml.SSMLSayAs{
            InterpretAs: SSMLTelephone,
            Text: "(510) 555 - 1234",
        },
    },
})

```

### Supports


#### [`SSMLText`](https://developer.amazon.com/docs/custom-skills/speech-synthesis-markup-language-ssml-reference.html#p)

Represents an `paragraph`

```golang
m := twiml.SSMLText{
    Text: "It is a good day.",
}
```

#### [`SSMLBreak`](https://developer.amazon.com/docs/custom-skills/speech-synthesis-markup-language-ssml-reference.html#break)

Represents an `break` and allows to fine tune a pause between speech.  

```golang
m := twiml.SSMLBreak{
    Strength: SSMLMedium, // or just "medium"
    Length: "850ms",
}
```

Strength can be one of `SSMLNone`, `SSMLXWeak`, `SSMLWeak`, `SSMLMedium`, `SMLStrong`, `SSMLXStrong`


#### [`SSMLSayAs`](https://developer.amazon.com/docs/custom-skills/speech-synthesis-markup-language-ssml-reference.html#say-as)

Describes how text should be interpreted.  

```golang
m := twiml.SSMLSayAs{
    InterpretAs: SSMLAddress
    Text: "124 Baker Street, Los Angeles, California",
}
```

InterpretAs can be one of `SSMLSpellOut`, `SSMLCharacters`, `SSMLOrdinal`, `SSMLDigits`, `SSMLFraction`, `SSMLUnit`, `SSMLDate`, `SSMLTime`, `SSMLTelephone`, `SSMLAddress`, `SSMLInterjection`, `SSMLExpletive`

#### [`SSMLEmphasis`](https://developer.amazon.com/docs/custom-skills/speech-synthesis-markup-language-ssml-reference.html#emphasis)

Emphasize the tagged words or phrases.

```golang
m := twiml.SSMLEmphasis{
    Level: SSMLStrong,
    Text: "It is a good day.",
}
```

Level can be one of `SSMLStrong`, `SSMLModerate`, `SSMLReduced`

#### [`SSMLProsody`](https://developer.amazon.com/docs/custom-skills/speech-synthesis-markup-language-ssml-reference.html#prosody)

Modifies the volume, pitch, and rate of the tagged speech.

```golang
m := twiml.SSMLProsody{
    Pitch: "+50%",
    Volume: SSMLXLoud,
    Text: "I am super loud high pitched voice",
}
```

#### [`SSMLEffect`](https://developer.amazon.com/docs/custom-skills/speech-synthesis-markup-language-ssml-reference.html#amazon-effect)

Applies Amazon-specific effects to the speech.

```golang
m := twiml.SSMLEffect{
    Name: SSMLEffectWhisper,
    Text: "There is a Ghost somewhere here",
}
```

Currently only `SSMLEffectWhisper` is supported.


## TwiMLString

TwiMLStrings is a simple string macro language allowing to define TwiML like verbs directly in a string.
It follows a simple format where a TwiML statement is encapsulated within brackets and the verb divided
by an | from it's optional content and attributes. The content can be encapsulated with single backward
quotes (`) to ensure easy compatibility with JSON. Attributes are separated by comma and follow the
format key:value.

### Format

```
{<verb>|<content>,<attr:key>:<attr:value>,...}
{<verb>|`<content>`,<attr:key>:<attr:value>,...}
```

### Examples

```golang
    ts := "{say|Hello World}"
    ts := "{say|Hello}{strong|World}{say|Hello World,voice:man}"
    ts := "{say|`Hello, World`,voice:female}"
```

### Advance Techniques

It tries to mimic as close as possible the TWIML verbs with their associated attributes however implements
assorted new verbs that allow quick use of advance SSML techniques such as saying an address or telephone number.

```golang 
ts := "{say|Call me back at}{telephone|(949)555-1234}"
```

### Supports

#### `play|<url>,loop:<int>`

```
{play|http://url-to-audio/file.mp3,loop:2}
```

#### `say|<text>,voice:<string>,language:<string>,loop:<int>`

```
{say|Hello World,voice:Brian}
```

Note: Polly voices are always enabled in TwiMLString.

#### `dial|<number>,Record:<int>`

```
{dial|Hello World,voice:Brian}
```

#### `dtfm|<digits>`

```
{dtfm|12345}
```

#### `hangup`

```
{hangup}
```

#### `pause|<time>`

```
{pause|2ms}
```

#### `strong|<text>,level:<string>`

```
{strong|Mark my words!,level:x-strong}
```

#### `p|<text>`

Speaks an paragraph with an natural pause at the end.

```
{p|It is called condensation}
```
#### `whisper|<text>,effect:<string>`

Whispers the text. In the future this may be renamed to effect.

```
{whisper|Maybe you can hear me}
```

#### SayAs support

All of the `SayAs` SSML definitions are defined and follow the principal format below:

```
{telephone|(555)123-4567} says it as phone number
{address|`17 Baker Street, Los Angeles`} says it as address
{spell-out|I AM AWESOME} spells out the text
```

All of the SayAs SSML identifier has been implemented.


## More Examples

### Processing a request from Twilio

The library contains helpers to bind incoming Twilio requests to a struct that includes all of the available info from the request.  Most initial requests from Twilio are of type `twiml.VoiceRequest`.  Other request types are possible as a result of callbacks you register in your response.  See the [GoDoc](https://godoc.org/github.com/BTBurke/twiml) for details.

```golang
func(w http.ResponseWriter, r *http.Request) {
    var vr twiml.VoiceRequest
    if err := twiml.Bind(&vr, r); err != nil {
        http.Error(w, http.StatusText(400), 400)
        return
    }
    fmt.Printf("Incoming call from %s", vr.From)
}
```

### Constructing a response using TwiML

Once you receive a request from the Twilio API, you construct a TwiML response to provide directions for how to deal with the call.  This library includes (most of) the allowable verbs and rules to validate that your response is constructed properly.

```golang
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


### Even more Examples

For a more detailed example of constructing a small TwiML response server, see my [Twilio Voice project](https://github.com/BTBurke/twilio-voice) which is a Google-voice clone that forwards calls to your number and handles transcribing voicemails.