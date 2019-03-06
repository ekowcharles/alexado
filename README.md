

# Alexa Data Objects

[![Build Status](https://travis-ci.com/ekowcharles/alexado.svg?branch=master)](https://travis-ci.com/ekowcharles/alexado) [![codecov](https://codecov.io/gh/ekowcharles/alexado/branch/master/graph/badge.svg)](https://codecov.io/gh/ekowcharles/alexado)

Go library that provides objects and basic behavior for the sending requests to and processing responses for Alexa.

## Usage

The following snippet show how to use this library to process request received from the Alexa platform:
```go
b, err := ioutil.ReadAll(r.Body) # 'r' is a pointer to http.Request
if err != nil {
  ... // handle error
}

areq := AlexaRequest{}
err = json.Unmarshal(b, &alexaRequest)
if err != nil {
  ... // handle error
}

// at this point 'areq' should have everything you need to access data in the Alexa request
```

The following demonstrates how you would use the alexado library to construct a response to be sent to the Alexa platform:
```go
ares := alexado.AlexaResponse{}
ares.Version = "1.0"

outputSpeech := alexado.OutputSpeech{}
outputSpeech.Type = alexado.SSML.String()
outputSpeech.SSML = "<speak>Alexa-do's got you covered!</speak>"
ares.Response.OutputSpeech = &outputSpeech

ares.Response.ShouldEndSession = true

responseBody, err = ares.toJSON()
if err != nil {
  ... // handle serialization error
}

w.WriteHeader(http.StatusOK) // is an http.ResponseWriter object
w.Header().Add("ContentType", "application/json")
io.WriteString(w, responseBody) // io is from the io/ioutil package
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
      "slots": [
        {
          "name": "speech",
          "value": "Alexa-do's got you covered!",
          "confirmationStatus": "NONE",
          "source": "USER"
        }
      ]
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

## Forum

[Join](https://join.slack.com/t/alexalibraries/shared_invite/enQtNTY3MDMyODU4ODk5LTBkMWNiNTVjYTY2MTJlMWMxY2M3YWI0NGFhY2Q3NzhhNWQ3ZDg5MWIyNzFmMTk0NTBlMzRiOGYyNTE4YjNlNzg) the discussion on slack.

## Contribution

Feel free to make a pull request. For bigger changes, create a issue first to discuss it.

# License

MIT License

Copyright (c) 2019 C A Boadu Jnr

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
