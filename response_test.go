package alexado

import "testing"

func TestOutputSpeechTypeString(t *testing.T) {
	var actual, expected string

	actual, expected = SSML.String(), "SSML"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = PlainText.String(), "PlainText"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}
}

func TestPlayBehaviorTypeString(t *testing.T) {
	var actual, expected string

	actual, expected = Enqueue.String(), "ENQUEUE"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = ReplaceAll.String(), "REPLACE_ALL"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = ReplaceEnqueued.String(), "REPLACE_ENQUEUED"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}
}

func TestCardTypeString(t *testing.T) {
	var actual, expected string

	actual, expected = Simple.String(), "Simple"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = Standard.String(), "Standard"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = LinkAccount.String(), "LinkAccount"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = AskForPermissionsConsent.String(), "AskForPermissionsConsent"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}
}

func TestToJSON(t *testing.T) {
	sa := Attributes{Key: "foo", Value: "bar"}
	osp := OutputSpeech{Type: SSML.String(), SSML: "Some good speech."}
	res := Response{OutputSpeech: &osp}
	req := AlexaResponse{Version: "1.0", Response: res, SessionAttributes: &sa}

	j, _ := req.ToJSON()

	e := string(`{"version":"1.0","sessionAttributes":{"key":"foo","value":"bar"},"response":{"outputSpeech":{"type":"SSML","ssml":"Some good speech."}}}`)

	if j != e {
		t.Errorf("JSON %s was not constructed properly.", j)
	}
}
