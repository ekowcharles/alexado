package alexa

type AlexaResponse struct {
	Version           string     `json:"version,omitempty"`
	SessionAttributes Attributes `json:"sessionAttributes,omitempty"`
	Response          Response   `json:"response,omitempty"`
}

type Response struct {
	OutputSpeech     OutputSpeech `json:"outputSpeech,omitempty"`
	Card             Card         `json:"card,omitempty"`
	Reprompt         Reprompt     `json:"reprompt,omitempty"`
	Directives       []Directives `json:"directives,omitempty"`
	ShouldEndSession bool         `json:"shouldEndSession,omitempty"`
}

type OutputSpeech struct {
	Type         string `json:"type,omitempty"`
	Text         string `json:"text,omitempty"`
	SSML         string `json:"ssml,omitempty"`
	PlayBehavior string `json:"playBehavior,omitempty"`
}

type OutputSpeechType int

const (
	SSML OutputSpeechType = 0
)

func (o OutputSpeechType) String() string {
	types := [...]string{
		"SSML",
	}

	return types[o]
}

type Card struct {
	Type  string `json:"type,omitempty"`
	Title string `json:"title,omitempty"`
	Text  string `json:"text,omitempty"`
	Image Image  `json:"image,omitempty"`
}

type Image struct {
	SmallImageURL string `json:"smallImageUrl,omitempty"`
	LargeImageURL string `json:"largeImageUrl,omitempty"`
}

type Reprompt struct {
	OutputSpeech OutputSpeech `json:"outputSpeech,omitempty"`
}

type Directives struct {
	Type string `json:"type,omitempty"`
}
