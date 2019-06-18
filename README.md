

# Alexa Data Objects

[![Build Status](https://travis-ci.com/ekowcharles/alexado.svg?branch=master)](https://travis-ci.com/ekowcharles/alexado) [![codecov](https://codecov.io/gh/ekowcharles/alexado/branch/master/graph/badge.svg)](https://codecov.io/gh/ekowcharles/alexado)
[![GoDoc](https://godoc.org/github.com/ekowcharles/alexado?status.svg)](https://godoc.org/github.com/ekowcharles/alexado)


Go library for serializing and deserializing Alexa platform objects.

## Usage

### Requests

#### Retrieving Alexa request data

The following snippet show how to use this library to process request received from the Alexa platform:
```go
b, err := ioutil.ReadAll(r.Body)  // 'r' is a pointer to http.Request
if err != nil {
  ...                             // handle error
}

areq := AlexaRequest{}
err = json.Unmarshal(b, &alexaRequest)
if err != nil {
  ...                            // handle error
}

// at this point 'areq' has everything you need to access data in the Alexa request
```

#### Accessing slots

The data in the request received from Amazon has dynamic content for the `slots` json node when it does include it. 

For example, given you receive the following request from the Alexa platform:
```json
  ...
  "intent": {
    "name": "NoteCreationIntent",
    "confirmationStatus": "NONE",
    "slots": {
      "day": {
        "name": "day",
        "value": "Friday",
        "confirmationStatus": "CONFIRMED",
        "source": "USER"
      },
      "speech": {
        "name": "note",
        "value": "It may actually rain",
        "confirmationStatus": "CONFIRMED",
        "source": "USER"
      }
    }
  }
  ...
```

You can extract the slots like so:
```go
alexaRequest.Request.Intent.Slots["day"].Name                  // 'day'
alexaRequest.Request.Intent.Slots["day"].Value                 // 'Friday'
alexaRequest.Request.Intent.Slots["day"].ConfirmationStatus    // 'CONFIRMED'
alexaRequest.Request.Intent.Slots["day"].Source                // 'USER'

alexaRequest.Request.Intent.Slots["note"].Name                 // 'note'
alexaRequest.Request.Intent.Slots["note"].Value                // 'It may actually rain'
alexaRequest.Request.Intent.Slots["note"].ConfirmationStatus   // 'CONFIRMED'
alexaRequest.Request.Intent.Slots["note"].Source               // 'USER'
```

Expect the following results for non-existent slots, (using 'missing' in this example)
```go
alexaRequest.Request.Intent.Slots["missing"]                    // '{   }', essentially zero value for a map
alexaRequest.Request.Intent.Slots["missing"].Name               // '', zero value for a string
alexaRequest.Request.Intent.Slots["missing"].Value              // ''
alexaRequest.Request.Intent.Slots["missing"].ConfirmationStatus // ''
alexaRequest.Request.Intent.Slots["missing"].Source             // ''
```

#### Handling time

Alexa uses the [RFC3339](https://tools.ietf.org/html/rfc3339) format for dates. Timestamps are automatically converted to this format in alexado for use.

#### Accessing request attributes

Given you receive the following in your request:
```json
...
"attributes" : {
  "context" : "update",
  "object" : "note",
  "id": "2690e1db-993f-4f29-a411-ad486a98d9e9"
}
...
```

You can retrieve the attributes like so:
```go
alexaRequest.Session.Attributes["context"]  // 'update'
alexaRequest.Session.Attributes["object"]   // 'note'
alexaRequest.Session.Attributes["id"]       // '2690e1db-993f-4f29-a411-ad486a98d9e9'
```

### Responses

#### Creating response object

The following demonstrates how you would use the alexado library to construct a response to be sent to the Alexa platform:
```go
ares := alexado.AlexaResponse{}
ares.Version = "1.0"

outputSpeech := alexado.OutputSpeech{}
outputSpeech.Type = alexado.SSML.String()
outputSpeech.SSML = "<speak>Alexa-do's got you covered!</speak>"
ares.Response.OutputSpeech = &outputSpeech

ares.Response.ShouldEndSession = true

responseBody, err = ares.ToJSON()
if err != nil {
  ...                                               // handle serialization error
}

w.WriteHeader(http.StatusOK)                        // is an http.ResponseWriter object
w.Header().Add("ContentType", "application/json")
io.WriteString(w, responseBody)                     // io is from the io/ioutil package
```
#### Setting session attributes

You can set the session attributes like so:
```go
ares := alexado.AlexaResponse{}
...
var s Attributes = make(map[string]string)
s["context"] = "update"
s["object"] = "note"
s["id"] = "2690e1db-993f-4f29-a411-ad486a98d9e9"
ares.SessionAttributes = s
...
```

## Samples

### Request from Alexa platform

```json
{
  "version": "1.0",
  "session": {
    "new": true,
    "sessionId": "amzn1.echo-api.session.[unique-value-here]",
    "application": {
      "applicationId": "amzn1.ask.skill.[unique-value-here]"
    },
    "attributes": {
      "key": "string value"
    },
    "user": {
      "userId": "amzn1.ask.account.[unique-value-here]",
      "accessToken": "Atza|AAAAAAAA...",
      "permissions": {
        "consentToken": "ZZZZZZZ..."
      }
    }
  },
  "context": {
    "System": {
      "device": {
        "deviceId": "string",
        "supportedInterfaces": {
          "AudioPlayer": {}
        }
      },
      "application": {
        "applicationId": "amzn1.ask.skill.[unique-value-here]"
      },
      "user": {
        "userId": "amzn1.ask.account.[unique-value-here]",
        "accessToken": "Atza|AAAAAAAA...",
        "permissions": {
          "consentToken": "ZZZZZZZ..."
        }
      },
      "apiEndpoint": "https://api.amazonalexa.com",
      "apiAccessToken": "AxThk..."
    },
    "AudioPlayer": {
      "playerActivity": "PLAYING",
      "token": "audioplayer-token",
      "offsetInMilliseconds": 0
    }
  },
  "request": {
    "type": "IntentRequest",
    "requestId": "amzn1.echo-api.request.[SomeIdentifier]",
    "timestamp": "2023-10-21T05:36:12Z",
    "locale": "en-US",
    "intent": {
      "name": "RecordIntent",
      "confirmationStatus": "NONE",
      "slots": {
        "speech": {
          "name": "speech",
          "value": "Alexa-do's got you covered!",
          "confirmationStatus": "NONE",
          "source": "USER"
        }
      }
    }
  }
}
```

### Response to Alexa platform

```json
{
  "version": "string",
  "sessionAttributes": {
    "key": "value"
  },
  "response": {
    "outputSpeech": {
      "type": "PlainText",
      "text": "Plain text string to speak",
      "playBehavior": "REPLACE_ENQUEUED"
    },
    "card": {
      "type": "Standard",
      "title": "Title of the card",
      "text": "Text content for a standard card",
      "image": {
        "smallImageUrl": "https://url-to-small-card-image...",
        "largeImageUrl": "https://url-to-large-card-image..."
      }
    },
    "reprompt": {
      "outputSpeech": {
        "type": "PlainText",
        "text": "Plain text string to speak",
        "playBehavior": "REPLACE_ENQUEUED"
      }
    },
    "directives": [
      {
        "type": "InterfaceName.Directive"
      }
    ],
    "shouldEndSession": true
  }
}
```

## Todo

Add support for [progressive response](https://developer.amazon.com/docs/custom-skills/send-the-user-a-progressive-response.html#directive-request).

```json
POST {context.System.apiEndpoint}/v1/directives HTTP/1.1
Authorization: Bearer {context.System.apiAccessToken}
Content-Type: application/json

{
  "header":{
    "requestId":"{request.requestId}"
  },
  "directive":{
    "type":"VoicePlayer.Speak",
    "speech":"This text is spoken while your skill processes the full response."
  }
}
```

Add some more tests for enumerations

## References

- [Request and Response JSON Reference](https://developer.amazon.com/docs/custom-skills/request-and-response-json-reference.html)

## Forum

[Join](https://join.slack.com/t/alexalibraries/shared_invite/enQtNTY3MDMyODU4ODk5LTBkMWNiNTVjYTY2MTJlMWMxY2M3YWI0NGFhY2Q3NzhhNWQ3ZDg5MWIyNzFmMTk0NTBlMzRiOGYyNTE4YjNlNzg) the discussion on slack.

## Contribution

Feel free to make a pull request. For bigger changes, create a issue first to discuss it. Be sure to add links to any useful references.

# License

MIT License

Copyright (c) 2019 C A Boadu Jnr

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
