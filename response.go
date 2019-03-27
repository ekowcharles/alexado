// Package alexado provides objects and basic behavior for the sending requests to and processing responses for Alexa
package alexado

import "encoding/json"

// AlexaResponse is the response to be sent to the Alexa platform
type AlexaResponse struct {
	Version           string     `json:"version,omitempty"`           // Version specifier for the response with the value to be defined as: "1.0"
	SessionAttributes Attributes `json:"sessionAttributes,omitempty"` // Map of key-value pairs to persist in the session
	Response          Response   `json:"response,omitempty"`          // Defines what to render to the user and whether to end the current session
}

// Attributes are key value pairs
type Attributes map[string]string

// Response defines what to render to the user and whether to end the current session
type Response struct {
	OutputSpeech     *OutputSpeech `json:"outputSpeech,omitempty"`     // Contains the type of output speech to render
	Card             *Card         `json:"card,omitempty"`             // Contains a card to render to the Amazon Alexa App
	Reprompt         *Reprompt     `json:"reprompt,omitempty"`         // Contains the outputSpeech to use if a re-prompt is necessary
	Directives       []Directive   `json:"directives,omitempty"`       // Specifies device-level actions to take using a particular interface, such as the AudioPlayer interface for streaming audio
	ShouldEndSession bool          `json:"shouldEndSession,omitempty"` // True meaning that the session should end after Alexa speaks the response, or false if the session should remain active. If not provided, defaults to true.
}

// OutputSpeech is used for setting both the outputSpeech and the reprompt properties
type OutputSpeech struct {
	Type         string `json:"type,omitempty"`         // Contains the type of output speech to render
	Text         string `json:"text,omitempty"`         // Contains the speech to render to the user
	SSML         string `json:"ssml,omitempty"`         // Contains text marked up with SSML to render to the user. Use this when type is  "SSML"
	PlayBehavior string `json:"playBehavior,omitempty"` // Determines the queuing and playback of this output speech
}

// OutputSpeechType is the type of output speech to render
type OutputSpeechType int

const (
	// SSML indicates that the output speech is text marked up with SSML
	SSML OutputSpeechType = iota
	// PlainText indicates that the output speech is defined as plain text
	PlainText
)

func (o OutputSpeechType) String() string {
	return [...]string{
		"SSML",
		"PlainText",
	}[o]
}

// PlayBehaviorType determines the queuing and playback of this output speech
type PlayBehaviorType int

const (
	// Enqueue adds this speech to the end of the queue. Do not interrupt Alexa's current speech. This is the default value for all skills that do not use the GameEngine interface
	Enqueue PlayBehaviorType = iota
	// ReplaceAll immediately begins playback of this speech, and replace any current and enqueued speech. This is the default value for all skills that use the GameEngine interface
	ReplaceAll
	// ReplaceEnqueued replaces all speech in the queue with this speech. Do not interrupt Alexa's current speech
	ReplaceEnqueued
)

func (p PlayBehaviorType) String() string {
	return [...]string{
		"ENQUEUE",
		"REPLACE_ALL",
		"REPLACE_ENQUEUED",
	}[p]
}

// Card can only be included when sending a response to a CanFulfillIntentRequest, LaunchRequest, IntentRequest, or InputHandlerEvent
type Card struct {
	Type    string `json:"type,omitempty"`    // Describes the type of card to render
	Title   string `json:"title,omitempty"`   // Contains the title of the card. (not applicable for cards of type LinkAccount).
	Text    string `json:"text,omitempty"`    // Contains the text content for a Standard card (not applicable for cards of type Simple or LinkAccount)
	Content string `json:"content,omitempty"` // Contains the text content for a Standard card (not applicable for cards of type Simple or LinkAccount)
	Image   *Image `json:"image,omitempty"`   // Specifies the URLs for the image to display on a Standard card. Only applicable for Standard cards.
}

// CardType describes the type of card to render
type CardType int

const (
	// Simple for when card contains a title and plain text content
	Simple CardType = iota
	// Standard for when card contains a title, text content, and an image to display
	Standard
	// LinkAccount for when card displays a link to an authorization URI that the user can use to link their Alexa account with a user in another system
	LinkAccount
	// AskForPermissionsConsent for when card asks the customer for consent to obtain specific customer information, such as Alexa lists or address information
	AskForPermissionsConsent
)

func (c CardType) String() string {
	return [...]string{
		"Simple",
		"Standard",
		"LinkAccount",
		"AskForPermissionsConsent",
	}[c]
}

// Image specifies the URLs for the image to display on a Standard card. Only applicable for Standard cards.
type Image struct {
	SmallImageURL string `json:"smallImageUrl,omitempty"` // Displayed on smaller screens
	LargeImageURL string `json:"largeImageUrl,omitempty"` // Displayed on larger screens
}

// Reprompt contains the outputSpeech to use if a re-prompt is necessary
type Reprompt struct {
	OutputSpeech OutputSpeech `json:"outputSpeech,omitempty"`
}

// Directive specifies device-level actions to take using a particular interface, such as the AudioPlayer interface for streaming audio
type Directive struct {
	Type string `json:"type,omitempty"`
}

// ToJSON converts the AlexaResponse object to json format
func (t AlexaResponse) ToJSON() (string, error) {
	toJSON, err := json.Marshal(t)

	return string(toJSON), err
}
