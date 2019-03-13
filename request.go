// Package alexado provides objects and basic behavior for the sending requests to and processing responses for Alexa
package alexado

import "time"

// AlexaRequest is expected request be sent from the Alexa plaform.
type AlexaRequest struct {
	Version string  `json:"version"` // Version specifier for the request with the value defined as: "1.0"
	Session Session `json:"session"` // Provides additional context associated with the request
	Context Context `json:"context"` // Provides your skill with information about the current state of the Alexa service and device at the time the request is sent to your service
	Request Request `json:"request"` // Provides the details of the user's request
}

// Session provides additional context associated with the request.
// Session holds data for standard request types CanFulfillIntentReqeuest, LaunchRequest, IntentRequest, and SessionEndedRequest. The GameEngine interface includes a session object also.
// Requests from interfaces such as AudioPlayer and PlaybackController are not sent in the context of a session, so they do not include the session object.
type Session struct {
	New         bool        `json:"new"`         // Indicates whether this is a new session. Returns true for a new session or false for an existing session.
	SessionID   string      `json:"sessionId"`   // Represents a unique identifier per a user's active session. A sessionId is consistent for multiple subsequent requests for a user and session. If the session ends for a user, then a new unique sessionId value is provided for subsequent requests for the same user.
	Application Application `json:"application"` // Contains an application ID. This is used to verify that the request was intended for your service.
	Attributes  Attributes  `json:"attributes"`  // Passed back to your skill on the next request
	User        User        `json:"user"`        // Describes the user making the request
}

// Context object provides your skill with information about the current state of the Alexa service and device at the time the request is sent to your service. This is included on all requests.
// For requests sent in the context of a session (CanFulfillIntentRequest, LaunchRequest and IntentRequest), the context object duplicates the user and application information that is also available in the session object.
type Context struct {
	System      System      `json:"System"`      // Provides information about the current state of the Alexa service and the device interacting with your skill
	AudioPlayer AudioPlayer `json:"AudioPlayer"` // Provides the current state for the AudioPlayer interface
	Viewport    Viewport    `json:"Viewport"`    // Describes the operating characteristics of the display device.
}

// Viewport describes the operating characteristics of the display device.
type Viewport struct {
	Experiences        []Experience `json:"experiences"`
	Shape              string       `json:"shape"`              // Shape of the viewport. RECTANGLE or ROUND.
	PixelWidth         int          `json:"pixelWidth"`         // Maximum viewport value
	PixelHeight        int          `json:"pixelHeight"`        // Maximum viewport value
	DPI                int          `json:"dpi"`                // Pixel density of the viewport
	CurrentPixelWidth  int          `json:"currentpixelWidth"`  // Viewport width that is currently in use
	CurrentPixelHeight int          `json:"currentpixelHeight"` // Viewport height that is currently in use
	Theme              string       `json:"theme"`              // Basic color scheme in use. LIGHT or DARK.
	Touch              []string     `json:"touch"`
	Keyboard           []string     `json:"keyboard"`
}

type Experience struct {
	ArcMinuteWidth  int  `json:"arcMinuteWidth"`
	ArcMinuteHeight int  `json:"arcMinuteHeight"`
	CanRotate       bool `json:"canRotate"`
	CanResize       bool `json:"canResize"`
}

// ShapeType is the shape of the viewport. RECTANGLE or ROUND.
type ShapeType int

const (
	// Rectangle for device with rectangular viewport
	Rectangle ShapeType = iota
	// Round is for device with round viewport
	Round
)

func (s ShapeType) String() string {
	return [...]string{
		"RECTANGLE",
		"ROUND",
	}[s]
}

// TouchType represents basic color scheme in use. LIGHT or DARK.
type TouchType int

const (
	Single TouchType = iota
)

func (t TouchType) String() string {
	return [...]string{
		"SINGLE",
	}[t]
}

type KeyboardType int

const (
	Direction KeyboardType = iota
)

func (k KeyboardType) String() string {
	return [...]string{
		"DIRECTION",
	}[k]
}

type ThemeType int

const (
	Light ThemeType = iota
	Dark
)

func (t ThemeType) String() string {
	return [...]string{
		"LIGHT",
		"DARK",
	}[t]
}

// System object provides information about the current state of the Alexa service and the device interacting with your skill.
type System struct {
	Device         Device      `json:"device"`         // Provides information about the device used to send the request
	Application    Application `json:"application"`    // Contains an application ID. This is used to verify that the request was intended for your service.
	User           User        `json:"user"`           // Describes the user making the request
	APIEndpoint    string      `json:"apiEndpoint"`    // References the correct base URI to refer to by region, for use with APIs such as the Device Location API and Progressive Response API.
	APIAccessToken string      `json:"apiAccessToken"` // contains a token that can be used to access Alexa-specific APIs
}

// Device provides information about the device used to send the request.
type Device struct {
	DeviceID            string              `json:"deviceId"`            // Uniquely identifies the device
	SupportedInterfaces SupportedInterfaces `json:"supportedInterfaces"` // Lists each interface that the device supports. For example, if supportedInterfaces includes AudioPlayer {}, then you know that the device supports streaming audio using the AudioPlayer interface.
}

// AudioPlayer provides the current state for the AudioPlayer interface.
// AudioPlayer is included on all customer-initiated requests (such as requests made by voice or using a remote control),
// but includes the details about the playback (token and offsetInMilliseconds) only when sent to a skill that was most recently playing audio.
type AudioPlayer struct {
	PlayerActivity       string `json:"playerActivity"`       // Indicates the last known state of audio playback
	Token                string `json:"token"`                // Represents the audio stream described by this AudioPlayer object
	OffsetInMilliseconds int    `json:"offsetInMilliseconds"` // Identifies a track's offset in milliseconds at the time the request was sent. This is 0 if the track is at the beginning.
}

// PlayerActivityType indicates the last known state of audio playback
type PlayerActivityType int

const (
	// Idle implies nothing was playing, no enqueued items.
	Idle PlayerActivityType = iota
	// Paused implies stream was paused.
	Paused
	// Playing implies stream was playing.
	Playing
	// BufferUnderrun implies buffer underrun
	BufferUnderrun
	// Finished implies stream was finished playing.
	Finished
	// Stopped implies stream was interrupted.
	Stopped
)

func (p PlayerActivityType) String() string {
	return [...]string{
		"IDLE",
		"PAUSED",
		"PLAYING",
		"BUFFER_UNDERRUN",
		"FINISHED",
		"STOPPED",
	}[p]
}

// SupportedInterfaces lists each interface that the device supports. For example, if supportedInterfaces includes AudioPlayer {},
// then you know that the device supports streaming audio using the AudioPlayer interface.
type SupportedInterfaces struct {
	AudioPlayer AudioPlayer `json:"audioPlayer"` // Lets you know that the device supports streaming audio using the AudioPlayer interface
}

// Application contains an application ID. This is used to verify that the request was intended for your service.
type Application struct {
	ApplicationID string `json:"applicationId"` // Represents the appliation ID for your skill.
}

// User describes the user making the request.
type User struct {
	UserID      string      `json:"userId"`      // Represents a unique identifier for the user who made the request. The length of this identifier can vary, but is never more than 255 characters. The userId is automatically generated when a user enables the skill in the Alexa app.
	AccessToken string      `json:"accessToken"` // Identifies the user in another system. This is only provided if the user has successfully linked their account.
	Permissions Permissions `json:"permissions"` // Contains a consentToken allowing the skill access to information that the customer has consented to provide, such as address information.
}

// Permissions contain a consentToken allowing the skill access to information that the customer has consented to provide, such as address information.
type Permissions struct {
	ContentToken string `json:"consentToken"` // Deprecated. Use the apiAccessToken available in the context object to determine the user's permissions.
}

// Request provides the details of the user's request. There are several different request types available.
type Request struct {
	RequestID string    `json:"requestId"`
	Timestamp time.Time `json:"timestamp"`
	Locale    string    `json:"locale"`
	Intent    Intent    `json:"intent"`
	Type      string    `json:"type"`
}

// Intent represents what user wants.
type Intent struct {
	Name               string          `json:"name"`
	ConfirmationStatus string          `json:"confirmationStatus"`
	Slots              map[string]Slot `json:"slots"`
}

// Slot represents user defined variables
type Slot struct {
	Name               string `json:"name"`
	Value              string `json:"value"`
	ConfirmationStatus string `json:"confirmationStatus"`
	Source             string `json:"source"`
}

//ConfirmationStatus is an enumeration indicating whether the user has explicitly confirmed or denied the value of this slot.
type ConfirmationStatusType int

const (
	// None indicates user has neither confirmed or denied the value of the slot.
	None ConfirmationStatusType = iota
	// Confirmed indicates user confirmed the value of the slot.
	Confirmed
	// Denied indicates user denied the value of the slot.
	Denied
)

func (c ConfirmationStatusType) String() string {
	return [...]string{
		"NONE",
		"CONFIRMED",
		"DENIED",
	}[c]
}

type SourceType int

const (
	UserSource SourceType = iota
)

func (s SourceType) String() string {
	return [...]string{
		"USER",
	}[s]
}

// RequestType describes types of requests to expect from the Alexa Platform.
type RequestType int

const (
	// LaunchRequest represents that a user made a request to an Alexa skill, but did not provide a specific intent.
	LaunchRequest RequestType = iota
	// CanFulfillIntentRequest represents a request made to skill to query whether the skill can understand and fulfill the intent request with detected slots, before actually asking the skill to take action.
	CanFulfillIntentRequest
	// SessionEndedRequest represents a request made to an Alexa skill to notify that a session was ended.
	SessionEndedRequest
	// IntentRequest represents a request made to a skill based on what the user wants to do.
	IntentRequest
)

// String returns request type as string.
func (r RequestType) String() string {
	return [...]string{
		"LaunchRequest",
		"CanFulfillIntentRequest",
		"SessionEndedRequest",
		"IntentRequest",
	}[r]
}

// IntentType is a higher level enumeration for Amazon built in intent types and custom intent type.
type IntentType int

// AmazonIntentType is an enumeration for Amazon intent types.
type AmazonIntentType IntentType

const (
	// AmazonCancelIntent lets the user cancel a transaction or task or lets the user completely exit the skill.
	AmazonCancelIntent AmazonIntentType = iota
	// AmazonFallbackIntent provides a fallback for user utterances that do not match any of your skill's intents.
	AmazonFallbackIntent
	// AmazonHelpIntent provides help about how to use the skill.
	AmazonHelpIntent
	// AmazonLoopOffIntent lets the user request that the skill turn off a loop or repeat mode, usually for audio skills that stream a playlist of tracks.
	AmazonLoopOffIntent
	// AmazonLoopOnIntent lets the user request that the skill turn on a loop or repeat mode, usually for audio skills that stream a playlist of tracks.
	AmazonLoopOnIntent
	// AmazonPauseIntent lets the user pause an action in progress, such as pausing a game or pausing an audio track that is playing.
	AmazonPauseIntent
	// AmazonPreviousIntent let the user go back to a previous item in a list.
	AmazonPreviousIntent
	// AmazonNextIntent lets the user navigate to the next item in a list.
	AmazonNextIntent
	// AmazonRepeatIntent lets the user request to repeat the last action.
	AmazonRepeatIntent
	// AmazonResumeIntent lets the user resume or continue an action.
	AmazonResumeIntent
	// AmazonSelectIntent lets users indicate that they want to select a particular item, such as an item on a list.
	AmazonSelectIntent
	// AmazonShuffleOffIntent lets the user request that the skill turn of a shuffle mode, usually for audio skills that stream a playlist of tracks.
	AmazonShuffleOffIntent
	// AmazonShuffleOnIntent lets the user request that the skill turn on a shuffle mode, usually for audio skills that stream a playlist of tracks.
	AmazonShuffleOnIntent
	// AmazonStartOverIntent lets the user request to restart an action, such as restarting a game, transaction, or audio track.
	AmazonStartOverIntent
	// AmazonStopIntent Either let the user stop an action (but remain in the skill) or lets the user completely exit the skill.
	AmazonStopIntent
	// AmazonYesIntent lets the user provide a positive response to a yes/no question for confirmation.
	AmazonYesIntent
	// AmazonNoIntent lets the user provide a negative response to a yes/no question for confirmation.
	AmazonNoIntent
)

// String returns intent type as string.
func (a AmazonIntentType) String() string {
	return [...]string{
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
	}[a]
}

// LocaleType represents the list of locales that Alexa supports
type LocaleType int

const (
	// DeDe de-DE German (DE)
	DeDe LocaleType = iota
	// EnAu en-AU English (AU)
	EnAu
	// EnCa en-CA English (CA)
	EnCa
	// EnGb en-GB English (UK)
	EnGb
	// EnIn en-IN English (IN)
	EnIn
	// EnUs en-US English (US)
	EnUs
	// EsEs es-ES Spanish (ES)
	EsEs
	// EsMx es-MX Spanish (MX)
	EsMx
	// FrCa fr-CA French (CA)
	FrCa
	// FrFr fr-FR French (FR)
	FrFr
	// ItIt it-IT Italian (IT)
	ItIt
	// JaJp ja-JP Japanese (JP)
	JaJp
)

func (l LocaleType) String() string {
	return [...]string{
		"de-DE",
		"en-AU",
		"en-CA",
		"en-GB",
		"en-IN",
		"en-US",
		"es-ES",
		"es-MX",
		"fr-CA",
		"fr-FR",
		"it-IT",
		"ja-JP",
	}[l]
}
