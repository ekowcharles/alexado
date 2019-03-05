package alexado

type AlexaRequest struct {
	Version string  `json:"version"`
	Session Session `json:"session"`
	Context Context `json:"context"`
	Request Request `json:"request"`
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

type ConfirmationStatus string

const (
	None ConfirmationStatus = "NONE"
)

type Source string

const (
	UserSource Source = "USER"
)

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
	AmazonCancelIntent           AmazonIntentType = 0
	AmazonFallbackIntent         AmazonIntentType = 1
	AmazonHelpIntent             AmazonIntentType = 2
	AmazonScrollUpIntent         AmazonIntentType = 3
	AmazonScrollLeftIntent       AmazonIntentType = 4
	AmazonScrollDownIntent       AmazonIntentType = 5
	AmazonScrollRightIntent      AmazonIntentType = 6
	AmazonPageUpIntent           AmazonIntentType = 7
	AmazonPageDownIntent         AmazonIntentType = 8
	AmazonMoreIntent             AmazonIntentType = 9
	AmazonNavigateHomeIntent     AmazonIntentType = 10
	AmazonNavigateSettingsIntent AmazonIntentType = 11
	AmazonLoopOffIntent          AmazonIntentType = 12
	AmazonLoopOnIntent           AmazonIntentType = 13
	AmazonNextIntent             AmazonIntentType = 14
	AmazonPauseIntent            AmazonIntentType = 15
	AmazonPreviousIntent         AmazonIntentType = 16
	AmazonNexIntent              AmazonIntentType = 17
	AmazonRepeatIntent           AmazonIntentType = 18
	AmazonResumeIntent           AmazonIntentType = 19
	AmazonSelectIntent           AmazonIntentType = 20
	AmazonShuffleOffIntent       AmazonIntentType = 21
	AmazonShuffleOnIntent        AmazonIntentType = 22
	AmazonStartOverIntent        AmazonIntentType = 23
	AmazonStopIntent             AmazonIntentType = 24
	AmazonYesIntent              AmazonIntentType = 25
	AmazonNoIntent               AmazonIntentType = 26
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
