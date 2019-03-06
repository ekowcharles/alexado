// Package alexado provides objects and basic behavior for the sending requests to and processing responses for Alexa
package alexado

// AlexaRequest is expected request be sent from the Alexa plaform.
type AlexaRequest struct {
	Version string  `json:"version"`
	Session Session `json:"session"`
	Context Context `json:"context"`
	Request Request `json:"request"`
}

// Session holds data for standard request types CanFulfillIntentReqeuest, LaunchRequest, IntentRequest, and SessionEndedRequest. The GameEngine interface includes a session object also.
// Requests from interfaces such as AudioPlayer and PlaybackController are not sent in the context of a session, so they do not include the session object.
type Session struct {
	New         bool        `json:"new"`
	SessionID   string      `json:"sessionId"`
	Application Application `json:"application"`
	Attributes  Attributes  `json:"attributes"`
	User        User        `json:"user"`
}

// Context object provides your skill with information about the current state of the Alexa service and device at the time the request is sent to your service. This is included on all requests.
// For requests sent in the context of a session (CanFulfillIntentRequest, LaunchRequest and IntentRequest), the context object duplicates the user and application information that is also available in the session object.
type Context struct {
	System      System      `json:"System"`
	AudioPlayer AudioPlayer `json:"AudioPlayer"`
}

// System object provides information about the current state of the Alexa service and the device interacting with your skill.
type System struct {
	Device         Device      `json:"device"`
	Application    Application `json:"application"`
	User           User        `json:"user"`
	APIEndpoint    string      `json:"apiEndpoint"`
	APIAccessToken string      `json:"apiAccessToken"`
}

// Device provides information about the device used to send the request.
type Device struct {
	DeviceID            string              `json:"deviceId"`
	SupportedInterfaces SupportedInterfaces `json:"supportedInterfaces"`
}

// AudioPlayer provides the current state for the AudioPlayer interface.
// AudioPlayer is included on all customer-initiated requests (such as requests made by voice or using a remote control),
// but includes the details about the playback (token and offsetInMilliseconds) only when sent to a skill that was most recently playing audio.
type AudioPlayer struct {
	PlayerActivity       string `json:"playerActivity"`
	Token                string `json:"token"`
	OffsetInMilliseconds int    `json:"offsetInMilliseconds"`
}

// SupportedInterfaces lists each interface that the device supports. For example, if supportedInterfaces includes AudioPlayer {},
// then you know that the device supports streaming audio using the AudioPlayer interface.
type SupportedInterfaces struct {
	AudioPlayer AudioPlayer `json:"audioPlayer"`
}

// Application contains an application ID. This is used to verify that the request was intended for your service.
type Application struct {
	ApplicationID string `json:"applicationId"`
}

// User describes the user making the request.
type User struct {
	UserID      string      `json:"userId"`
	AccessToken string      `json:"accessToken"`
	Permissions Permissions `json:"permissions"`
}

// Permissions contain a consentToken allowing the skill access to information that the customer has consented to provide, such as address information.
type Permissions struct {
	ContentToken string `json:"consentToken"`
}

// Request provides the details of the user's request. There are several different request types available.
type Request struct {
	RequestID string `json:"requestId"`
	Timestamp string `json:"timestamp"`
	Locale    string `json:"locale"`
	Intent    Intent `json:"intent"`
	Type      string `json:"type"`
}

// Intent represents what user wants.
type Intent struct {
	Name               string             `json:"name"`
	ConfirmationStatus ConfirmationStatus `json:"confirmationStatus"`
	Slots              []Slot             `json:"slots"`
}

type Slot struct {
	Name               string             `json:"name"`
	Value              string             `json:"value"`
	ConfirmationStatus ConfirmationStatus `json:"confirmationStatus"`
	Source             Source             `json:"source"`
}

//ConfirmationStatus is an enumeration indicating whether the user has explicitly confirmed or denied the value of this slot.
type ConfirmationStatus string

const (
	// None indicates user has neither confirmed or denied the value of the slot.
	None ConfirmationStatus = "NONE"
	// Confirmed indicates user confirmed the value of the slot.
	Confirmed ConfirmationStatus = "CONFIRMED"
	// Denied indicates user denied the value of the slot.
	Denied ConfirmationStatus = "DENIED"
)

type Source string

const (
	UserSource Source = "USER"
)

// RequestType describes types of requests to expect from the Alexa Platform.
type RequestType int

const (
	// LaunchRequest represents that a user made a request to an Alexa skill, but did not provide a specific intent.
	LaunchRequest RequestType = 0
	// CanFulfillIntentRequest represents a request made to skill to query whether the skill can understand and fulfill the intent request with detected slots, before actually asking the skill to take action.
	CanFulfillIntentRequest RequestType = 1
	// SessionEndedRequest represents a request made to an Alexa skill to notify that a session was ended.
	SessionEndedRequest RequestType = 2
	// IntentRequest represents a request made to a skill based on what the user wants to do.
	IntentRequest RequestType = 3
)

// String returns request type as string.
func (r RequestType) String() string {
	types := [...]string{
		"LaunchRequest",
		"CanFulfillIntentRequest",
		"SessionEndedRequest",
		"IntentRequest",
	}

	return types[r]
}

// IsLaunchRequest returns true when the intent type is LaunchRequest.
func (r Request) IsLaunchRequest() bool {
	if r.Type == LaunchRequest.String() {
		return true
	}

	return false
}

// IsCanFulfillIntentRequest returns true when the intent type is CanFulfillIntentRequest.
func (r Request) IsCanFulfillIntentRequest() bool {
	if r.Type == CanFulfillIntentRequest.String() {
		return true
	}

	return false
}

// IsSessionEndedRequest returns true when the inten type is SessionEndedRequest.
func (r Request) IsSessionEndedRequest() bool {
	if r.Type == SessionEndedRequest.String() {
		return true
	}

	return false
}

// IsIntentRequest returns true when the inten type is IntentRequest.
func (r Request) IsIntentRequest() bool {
	if r.Type == IntentRequest.String() {
		return true
	}

	return false
}

// IntentType is a higher level enumeration for Amazon built in intent types and custom intent type.
type IntentType int

// AmazonIntentType is an enumeration for Amazon intent types.
type AmazonIntentType IntentType

const (
	// AmazonCancelIntent lets the user cancel a transaction or task or lets the user completely exit the skill.
	AmazonCancelIntent AmazonIntentType = 0
	// AmazonFallbackIntent provides a fallback for user utterances that do not match any of your skill's intents.
	AmazonFallbackIntent AmazonIntentType = 1
	// AmazonHelpIntent provides help about how to use the skill.
	AmazonHelpIntent AmazonIntentType = 2
	// AmazonLoopOffIntent lets the user request that the skill turn off a loop or repeat mode, usually for audio skills that stream a playlist of tracks.
	AmazonLoopOffIntent AmazonIntentType = 3
	// AmazonLoopOnIntent lets the user request that the skill turn on a loop or repeat mode, usually for audio skills that stream a playlist of tracks.
	AmazonLoopOnIntent AmazonIntentType = 4
	// AmazonPauseIntent lets the user pause an action in progress, such as pausing a game or pausing an audio track that is playing.
	AmazonPauseIntent AmazonIntentType = 5
	// AmazonPreviousIntent let the user go back to a previous item in a list.
	AmazonPreviousIntent AmazonIntentType = 6
	// AmazonNextIntent lets the user navigate to the next item in a list.
	AmazonNextIntent AmazonIntentType = 7
	// AmazonRepeatIntent lets the user request to repeat the last action.
	AmazonRepeatIntent AmazonIntentType = 8
	// AmazonResumeIntent lets the user resume or continue an action.
	AmazonResumeIntent AmazonIntentType = 9
	// AmazonSelectIntent lets users indicate that they want to select a particular item, such as an item on a list.
	AmazonSelectIntent AmazonIntentType = 10
	// AmazonShuffleOffIntent lets the user request that the skill turn of a shuffle mode, usually for audio skills that stream a playlist of tracks.
	AmazonShuffleOffIntent AmazonIntentType = 11
	// AmazonShuffleOnIntent lets the user request that the skill turn on a shuffle mode, usually for audio skills that stream a playlist of tracks.
	AmazonShuffleOnIntent AmazonIntentType = 12
	// AmazonStartOverIntent lets the user request to restart an action, such as restarting a game, transaction, or audio track.
	AmazonStartOverIntent AmazonIntentType = 13
	// AmazonStopIntent Either let the user stop an action (but remain in the skill) or lets the user completely exit the skill.
	AmazonStopIntent AmazonIntentType = 14
	// AmazonYesIntent lets the user provide a positive response to a yes/no question for confirmation.
	AmazonYesIntent AmazonIntentType = 15
	// AmazonNoIntent lets the user provide a negative response to a yes/no question for confirmation.
	AmazonNoIntent AmazonIntentType = 16
)

// String returns intent type as string.
func (a AmazonIntentType) String() string {
	types := [...]string{
		"AMAZON.CancelIntent",
		"AMAZON.FallbackIntent",
		"AMAZON.HelpIntent",
		"AMAZON.LoopOffIntent",
		"AMAZON.LoopOnIntent",
		"AMAZON.PauseIntent",
		"AMAZON.PreviousIntent",
		"AMAZON.NextIntent",
		"AMAZON.RepeatIntent",
		"AMAZON.ResumeIntent",
		"AMAZON.SelectIntent",
		"AMAZON.ShuffleOffIntent",
		"AMAZON.ShuffleOnIntent",
		"AMAZON.StartOverIntent",
		"AMAZON.StopIntent",
		"AMAZON.YesIntent",
		"AMAZON.NoIntent",
	}

	return types[a]
}
