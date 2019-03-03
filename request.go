package alexa

import (
	"encoding/json"
)

type AlexaRequest struct {
	Version string  `json:"version,omitempty"`
	Session Session `json:"session"`
	Context Context `json:"context"`
	Request Request `json:"request"`
}

func (t AlexaRequest) toJSON() string {
	toJSON, _ := json.Marshal(t)

	return string(toJSON)
}

type Session struct {
	New         bool        `json:"new"`
	SessionID   string      `json:"sessionId"`
	Application Application `json:"application"`
	Attributes  Attributes  `json:"attributes"`
	User        User        `json:"user"`
}

type Context struct {
	System      System      `json:"System"`
	AudioPlayer AudioPlayer `json:"AudioPlayer"`
}

type System struct {
	Device         Device      `json:"device"`
	Application    Application `json:"application"`
	User           User        `json:"user"`
	APIEndpoint    string      `json:"apiEndpoint"`
	APIAccessToken string      `json:"apiAccessToken"`
}

type Attributes struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Device struct {
	DeviceID            string              `json:"deviceId"`
	SupportedInterfaces SupportedInterfaces `json:"supportedInterfaces"`
}

type AudioPlayer struct {
	PlayerActivity       string `json:"playerActivity"`
	Token                string `json:"token"`
	OffsetInMilliseconds int    `json:"offsetInMilliseconds"`
}

type SupportedInterfaces struct {
	AudioPlayer AudioPlayer `json:"audioPlayer"`
}

type Application struct {
	ApplicationID string `json:"applicationId"`
}

type User struct {
	UserID      string      `json:"userId"`
	AccessToken string      `json:"accessToken"`
	Permissions Permissions `json:"permissions"`
}

type Permissions struct {
	ContentToken string `json:"consentToken"`
}

type Request struct {
	RequestID string `json:"requestId"`
	Timestamp string `json:"timestamp"`
	Locale    string `json:"locale"`
	Intent    Intent `json:"intent"`
	Type      string `json:"type"`
}

type Intent struct {
	Name               string `json:"name"`
	ConfirmationStatus string `json:"confirmationStatus"`
	Slots              []Slot `json:"slots"`
}

type Slot struct {
	Name               string             `json:"name"`
	Value              string             `json:"value"`
	ConfirmationStatus ConfirmationStatus `json:"confirmationStatus"`
	Source             Source             `json:"source"`
}

type ConfirmationStatus int

const (
	None ConfirmationStatus = 0
)

func (c ConfirmationStatus) String() string {
	types := [...]string{
		"NONE",
	}

	return types[c]
}

type Source int

const (
	UserSource Source = 0
)

func (s Source) String() string {
	types := [...]string{
		"USER",
	}

	return types[s]
}

type RequestType int

const (
	LaunchRequest       RequestType = 0
	SessionEndedRequest RequestType = 1
	IntentRequest       RequestType = 2
)

func (r RequestType) String() string {
	types := [...]string{
		"LaunchRequest",
		"SessionEndedRequest",
		"IntentRequest"}

	return types[r]
}

func (r Request) IsLaunchRequest() bool {
	if r.Type == LaunchRequest.String() {
		return true
	}

	return false
}

func (r Request) IsSessionEndedRequest() bool {
	if r.Type == SessionEndedRequest.String() {
		return true
	}

	return false
}

func (r Request) IsIntentRequest() bool {
	if r.Type == IntentRequest.String() {
		return true
	}

	return false
}

type IntentType int

type AmazonIntentType IntentType

const (
	AmazonCancelIntent           AmazonIntentType = 1
	AmazonFallbackIntent         AmazonIntentType = 2
	AmazonHelpIntent             AmazonIntentType = 3
	AmazonScrollUpIntent         AmazonIntentType = 4
	AmazonScrollLeftIntent       AmazonIntentType = 5
	AmazonScrollDownIntent       AmazonIntentType = 6
	AmazonScrollRightIntent      AmazonIntentType = 7
	AmazonPageUpIntent           AmazonIntentType = 8
	AmazonPageDownIntent         AmazonIntentType = 9
	AmazonMoreIntent             AmazonIntentType = 10
	AmazonNavigateHomeIntent     AmazonIntentType = 11
	AmazonNavigateSettingsIntent AmazonIntentType = 12
	AmazonLoopOffIntent          AmazonIntentType = 13
	AmazonLoopOnIntent           AmazonIntentType = 14
	AmazonNextIntent             AmazonIntentType = 15
	AmazonPauseIntent            AmazonIntentType = 16
	AmazonPreviousIntent         AmazonIntentType = 17
	AmazonNexIntent              AmazonIntentType = 18
	AmazonRepeatIntent           AmazonIntentType = 19
	AmazonResumeIntent           AmazonIntentType = 20
	AmazonSelectIntent           AmazonIntentType = 21
	AmazonShuffleOffIntent       AmazonIntentType = 22
	AmazonShuffleOnIntent        AmazonIntentType = 23
	AmazonStartOverIntent        AmazonIntentType = 24
	AmazonStopIntent             AmazonIntentType = 25
	AmazonYesIntent              AmazonIntentType = 26
	AmazonNoIntent               AmazonIntentType = 27
)

func (a AmazonIntentType) String() string {
	types := [...]string{
		"AMAZON.CancelIntent",
		"AMAZON.FallbackIntent",
		"AMAZON.HelpIntent",
		"AMAZON.ScrollUpIntent",
		"AMAZON.ScrollLeftIntent",
		"AMAZON.ScrollDownIntent",
		"AMAZON.ScrollRightIntent",
		"AMAZON.PageUpIntent",
		"AMAZON.PageDownIntent",
		"AMAZON.MoreIntent",
		"AMAZON.NavigateHomeIntent",
		"AMAZON.NavigateSettingsIntent",
		"AMAZON.LoopOffIntent",
		"AMAZON.LoopOnIntent",
		"AMAZON.NextIntent",
		"AMAZON.PauseIntent",
		"AMAZON.PreviousIntent",
		"AMAZON.NexIntent",
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
