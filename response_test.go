package alexado

import "testing"

func TestToJSON(t *testing.T) {
	sa := Attributes{Key: "foo", Value: "bar"}
	osp := OutputSpeech{Type: SSML.String(), SSML: "Some good speech."}
	res := Response{OutputSpeech: &osp}
	req := AlexaResponse{Version: "1.0", Response: res, SessionAttributes: &sa}

	j := req.toJSON()

	e := string(`{"version":"1.0","sessionAttributes":{"key":"foo","value":"bar"},"response":{"outputSpeech":{"type":"SSML","ssml":"Some good speech."}}}`)

	if j != e {
		t.Errorf("JSON %s was not constructed properly.", j)
	}
}
