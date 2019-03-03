package alexa

type AlexaResponse struct {
	Version           string     `json:"version"`
	SessionAttributes Attributes `json:"sessionAttributes"`
	Response          Response   `json:"response"`
}

type Response struct {
	OutputSpeech     OutputSpeech `json:"outputSpeech"`
	Card             Card         `json:"card"`
	Reprompt         Reprompt     `json:"reprompt"`
	Directives       []Directive  `json:"directives"`
	ShouldEndSession bool         `json:"shouldEndSession"`
}

type OutputSpeech struct {
	Type         string `json:"type"`
	Text         string `json:"text"`
	SSML         string `json:"ssml"`
	PlayBehavior string `json:"playBehavior"`
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
	Type  string `json:"type"`
	Title string `json:"title"`
	Text  string `json:"text"`
	Image Image  `json:"image"`
}

type Image struct {
	SmallImageURL string `json:"smallImageUrl"`
	LargeImageURL string `json:"largeImageUrl"`
}

type Reprompt struct {
	OutputSpeech OutputSpeech `json:"outputSpeech"`
}

type Directive struct {
	Type string `json:"type"`
}
