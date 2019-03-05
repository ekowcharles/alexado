package alexado

import "testing"

func TestRequestTypeString(t *testing.T) {
	var actual, expected string

	actual, expected = LaunchRequest.String(), "LaunchRequest"
	if actual != expected {
		t.Errorf("%s != %s", actual, expected)
	}

	actual, expected = SessionEndedRequest.String(), "SessionEndedRequest"
	if actual != expected {
		t.Errorf("%s != %s", actual, expected)
	}

	actual, expected = IntentRequest.String(), "IntentRequest"
	if actual != expected {
		t.Errorf("%s != %s", actual, expected)
	}
}

func TestTrueIsLaunchRequest(t *testing.T) {
	r := Request{Type: LaunchRequest.String()}

	if !r.IsLaunchRequest() {
		t.Errorf("%s is not a LaunchRequest", r.Type)
	}
}

func TestFalseIsLaunchRequest(t *testing.T) {
	r := Request{Type: "SomeRequest"}

	if r.IsLaunchRequest() {
		t.Errorf("%s is not a LaunchRequest", r.Type)
	}
}

func TestTrueIsSessionEndedRequest(t *testing.T) {
	r := Request{Type: SessionEndedRequest.String()}

	if !r.IsSessionEndedRequest() {
		t.Errorf("%s is not a SessionEndedRequest", r.Type)
	}
}

func TestFalseIsSessionEndedRequest(t *testing.T) {
	r := Request{Type: "SomeRequest"}

	if r.IsSessionEndedRequest() {
		t.Errorf("%s is not a SessionEndedRequest", r.Type)
	}
}

func TestTrueIsIntentRequest(t *testing.T) {
	r := Request{Type: IntentRequest.String()}

	if !r.IsIntentRequest() {
		t.Errorf("%s is not a IntentRequest", r.Type)
	}
}

func TestFalseIsIntentRequest(t *testing.T) {
	r := Request{Type: "SomeRequest"}

	if r.IsIntentRequest() {
		t.Errorf("%s is not a IntentRequest", r.Type)
	}
}

func TestAmazonIntentTypeString(t *testing.T) {
	var actual, expected string

	actual, expected = AmazonCancelIntent.String(), "AMAZON.CancelIntent"
	if actual != expected {
		t.Errorf("%s != %s", actual, expected)
	}

	actual, expected = AmazonFallbackIntent.String(), "AMAZON.FallbackIntent"
	if actual != expected {
		t.Errorf("%s != %s", actual, expected)
	}

	actual, expected = AmazonHelpIntent.String(), "AMAZON.HelpIntent"
	if actual != expected {
		t.Errorf("%s != %s", actual, expected)
	}

	actual, expected = AmazonScrollUpIntent.String(), "AMAZON.ScrollUpIntent"
	if actual != expected {
		t.Errorf("%s != %s", actual, expected)
	}

	actual, expected = AmazonScrollLeftIntent.String(), "AMAZON.ScrollLeftIntent"
	if actual != expected {
		t.Errorf("%s != %s", actual, expected)
	}

	actual, expected = AmazonScrollDownIntent.String(), "AMAZON.ScrollDownIntent"
	if actual != expected {
		t.Errorf("%s != %s", actual, expected)
	}

	actual, expected = AmazonScrollRightIntent.String(), "AMAZON.ScrollRightIntent"
	if actual != expected {
		t.Errorf("%s != %s", actual, expected)
	}

	actual, expected = AmazonPageUpIntent.String(), "AMAZON.PageUpIntent"
	if actual != expected {
		t.Errorf("%s != %s", actual, expected)
	}

	actual, expected = AmazonPageDownIntent.String(), "AMAZON.PageDownIntent"
	if actual != expected {
		t.Errorf("%s != %s", actual, expected)
	}

	actual, expected = AmazonMoreIntent.String(), "AMAZON.MoreIntent"
	if actual != expected {
		t.Errorf("%s != %s", actual, expected)
	}

	actual, expected = AmazonNavigateHomeIntent.String(), "AMAZON.NavigateHomeIntent"
	if actual != expected {
		t.Errorf("%s != %s", actual, expected)
	}

	actual, expected = AmazonNavigateSettingsIntent.String(), "AMAZON.NavigateSettingsIntent"
	if actual != expected {
		t.Errorf("%s != %s", actual, expected)
	}

	actual, expected = AmazonLoopOffIntent.String(), "AMAZON.LoopOffIntent"
	if actual != expected {
		t.Errorf("%s != %s", actual, expected)
	}

	actual, expected = AmazonLoopOnIntent.String(), "AMAZON.LoopOnIntent"
	if actual != expected {
		t.Errorf("%s != %s", actual, expected)
	}

	actual, expected = AmazonNextIntent.String(), "AMAZON.NextIntent"
	if actual != expected {
		t.Errorf("%s != %s", actual, expected)
	}

	actual, expected = AmazonPauseIntent.String(), "AMAZON.PauseIntent"
	if actual != expected {
		t.Errorf("%s != %s", actual, expected)
	}

	actual, expected = AmazonPreviousIntent.String(), "AMAZON.PreviousIntent"
	if actual != expected {
		t.Errorf("%s != %s", actual, expected)
	}

	actual, expected = AmazonNexIntent.String(), "AMAZON.NexIntent"
	if actual != expected {
		t.Errorf("%s != %s", actual, expected)
	}

	actual, expected = AmazonRepeatIntent.String(), "AMAZON.RepeatIntent"
	if actual != expected {
		t.Errorf("%s != %s", actual, expected)
	}

	actual, expected = AmazonResumeIntent.String(), "AMAZON.ResumeIntent"
	if actual != expected {
		t.Errorf("%s != %s", actual, expected)
	}

	actual, expected = AmazonSelectIntent.String(), "AMAZON.SelectIntent"
	if actual != expected {
		t.Errorf("%s != %s", actual, expected)
	}

	actual, expected = AmazonShuffleOffIntent.String(), "AMAZON.ShuffleOffIntent"
	if actual != expected {
		t.Errorf("%s != %s", actual, expected)
	}

	actual, expected = AmazonShuffleOnIntent.String(), "AMAZON.ShuffleOnIntent"
	if actual != expected {
		t.Errorf("%s != %s", actual, expected)
	}

	actual, expected = AmazonStartOverIntent.String(), "AMAZON.StartOverIntent"
	if actual != expected {
		t.Errorf("%s != %s", actual, expected)
	}

	actual, expected = AmazonStopIntent.String(), "AMAZON.StopIntent"
	if actual != expected {
		t.Errorf("%s != %s", actual, expected)
	}

	actual, expected = AmazonYesIntent.String(), "AMAZON.YesIntent"
	if actual != expected {
		t.Errorf("%s != %s", actual, expected)
	}

	actual, expected = AmazonNoIntent.String(), "AMAZON.NoIntent"
	if actual != expected {
		t.Errorf("%s != %s", actual, expected)
	}
}
