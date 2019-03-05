

# Alexa Data Objects

[![Build Status](https://travis-ci.com/ekowcharles/alexado.svg?branch=master)](https://travis-ci.com/ekowcharles/alexado) [![codecov](https://codecov.io/gh/ekowcharles/alexado/branch/master/graph/badge.svg)](https://codecov.io/gh/ekowcharles/alexado)

Go library that provides objects and basic behaviors for the sending requests to and processing responses for Alexa.

## Usage

The following snippet show how to use this library to process request received from the Alexa platform:
```go
b, err := ioutil.ReadAll(r.Body) # 'r' is a pointer to http.Request
if err != nil {
	... # handle error
}

ar := alexa.AlexaRequest{}
err = json.Unmarshal(b, &alexaRequest)
if err != nil {
	... # handle error
}

# at this point 'ar' should have everything you need to access data in the Alexa request
```

The following demonstrates how you would use the alexado library to construct a response to be sent to the Alexa platform:
```go
alexaResponse := alexa.AlexaResponse{}
alexaResponse.Version = "1.0"

outputSpeech := alexa.OutputSpeech{}
outputSpeech.Type = alexa.SSML.String()
outputSpeech.SSML = "<speak>Alexa-do's got you covered!</speak>"

alexaResponse.Response.ShouldEndSession = true

alexaResponse.Response.OutputSpeech = &outputSpeech

w.WriteHeader(http.StatusOK) # is an http.ResponseWriter object
w.Header().Add("ContentType", "application/json")
io.WriteString(w, string(responseBody)) # io is from the io/ioutil package
```

## Samples

### Request from Alexa platform

```json
{
	"version": "1.0",
	"session": {
		"new": false,
		"sessionId": "amzn1.echo-api.session.[SomeIdentifier]",
		"application": {
			"applicationId": "amzn1.ask.skill.[SomeIdentifier]"
		},
		"user": {
			"userId": "amzn1.ask.account.[SomeRandomString]"
		}
	},
	"context": {
		"System": {
			"application": {
				"applicationId": "amzn1.ask.skill.[SomeIdentifier]"
			},
			"user": {
				"userId": "amzn1.ask.account.[SomeRandomString]"
			},
			"device": {
				"deviceId": "amzn1.ask.device.[SomeRandomString]",
				"supportedInterfaces": {}
			},
			"apiEndpoint": "https://api.amazonalexa.com",
			"apiAccessToken": "[ReallyLongRandomString]"
		},
		"Viewport": {
			"experiences": [{
				"arcMinuteWidth": 246,
				"arcMinuteHeight": 144,
				"canRotate": false,
				"canResize": false
			}],
			"shape": "RECTANGLE",
			"pixelWidth": 1024,
			"pixelHeight": 600,
			"dpi": 160,
			"currentPixelWidth": 1024,
			"currentPixelHeight": 600,
			"touch": ["SINGLE"]
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
			"slots": [{
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
    "version": "1.0",
    "response": {
        "outputSpeech": {
            "type": "SSML",
            "ssml": "<speak>Alexa-do's got you covered!</speak>"
        }
    }
}
```


# License

MIT License

Copyright (c) 2019 C A Boadu Jnr

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
