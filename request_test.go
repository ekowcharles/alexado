package alexado

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestShapeTypeString(t *testing.T) {
	var actual, expected string

	actual, expected = Rectangle.String(), "RECTANGLE"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = Round.String(), "ROUND"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}
}

func TestTouchTypeString(t *testing.T) {
	var actual, expected string

	actual, expected = Single.String(), "SINGLE"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}
}

func TestKeyboardTypeString(t *testing.T) {
	var actual, expected string

	actual, expected = Direction.String(), "DIRECTION"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}
}
func TestThemeTypeString(t *testing.T) {
	var actual, expected string

	actual, expected = Light.String(), "LIGHT"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = Dark.String(), "DARK"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}
}

func TestPlayerActivityTypeString(t *testing.T) {
	var actual, expected string

	actual, expected = Idle.String(), "IDLE"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = Paused.String(), "PAUSED"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = Playing.String(), "PLAYING"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = BufferUnderrun.String(), "BUFFER_UNDERRUN"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = Finished.String(), "FINISHED"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = Stopped.String(), "STOPPED"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}
}

func TestConfirmationStatusTypeString(t *testing.T) {
	var actual, expected string

	actual, expected = None.String(), "NONE"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = Confirmed.String(), "CONFIRMED"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = Denied.String(), "DENIED"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}
}

func TestSourceTypeString(t *testing.T) {
	var actual, expected string

	actual, expected = UserSource.String(), "USER"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}
}

func TestRequestTypeString(t *testing.T) {
	var actual, expected string

	actual, expected = LaunchRequest.String(), "LaunchRequest"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = CanFulfillIntentRequest.String(), "CanFulfillIntentRequest"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = SessionEndedRequest.String(), "SessionEndedRequest"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = IntentRequest.String(), "IntentRequest"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}
}

func TestAmazonIntentTypeString(t *testing.T) {
	var actual, expected string

	actual, expected = AmazonCancelIntent.String(), "AMAZON.CancelIntent"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = AmazonFallbackIntent.String(), "AMAZON.FallbackIntent"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = AmazonHelpIntent.String(), "AMAZON.HelpIntent"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = AmazonLoopOffIntent.String(), "AMAZON.LoopOffIntent"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = AmazonLoopOnIntent.String(), "AMAZON.LoopOnIntent"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = AmazonPauseIntent.String(), "AMAZON.PauseIntent"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = AmazonPreviousIntent.String(), "AMAZON.PreviousIntent"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = AmazonNextIntent.String(), "AMAZON.NextIntent"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = AmazonRepeatIntent.String(), "AMAZON.RepeatIntent"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = AmazonResumeIntent.String(), "AMAZON.ResumeIntent"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = AmazonSelectIntent.String(), "AMAZON.SelectIntent"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = AmazonShuffleOffIntent.String(), "AMAZON.ShuffleOffIntent"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = AmazonShuffleOnIntent.String(), "AMAZON.ShuffleOnIntent"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = AmazonStartOverIntent.String(), "AMAZON.StartOverIntent"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = AmazonStopIntent.String(), "AMAZON.StopIntent"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = AmazonYesIntent.String(), "AMAZON.YesIntent"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = AmazonNoIntent.String(), "AMAZON.NoIntent"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}
}

func TestLocaleTypeString(t *testing.T) {
	var actual, expected string

	actual, expected = DeDe.String(), "de-DE"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = EnAu.String(), "en-AU"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = EnCa.String(), "en-CA"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = EnGb.String(), "en-GB"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = EnIn.String(), "en-IN"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = EnUs.String(), "en-US"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = EsEs.String(), "es-ES"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = EsMx.String(), "es-MX"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = FrCa.String(), "fr-CA"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = FrFr.String(), "fr-FR"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = ItIt.String(), "it-IT"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = JaJp.String(), "ja-JP"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}
}

// The quickest way to test this is to compare the loaded json sample with the json result when the AlexaRequest is marshalled
// While this tests that the json result is the same as the json input, it does not test that the proper resource dependencies
// or nesting exist on the AlexaRequest Object hence this long seemingly convoluted approach to testing
func TestRequestUnmarshallsCorrectly(t *testing.T) {
	jsonFile, _ := os.Open("sample/request.json")
	defer jsonFile.Close()

	content, _ := ioutil.ReadAll(jsonFile)

	var alexaRequest = new(AlexaRequest)

	json.Unmarshal(content, alexaRequest)

	var actual, expected interface{}

	actual, expected = alexaRequest.Version, "1.0"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	session := alexaRequest.Session
	actual, expected = session.New, false
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = session.SessionID, "amzn1.echo-api.session.fcb76599-689e-401a-b305-d5556944f933"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = session.Application.ApplicationID, "amzn1.ask.skill.2604e1db-063f-4f29-a411-6eba27b9dc34"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = session.User.UserID, "amzn1.ask.account.userid"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	system := alexaRequest.Context.System
	actual, expected = system.Application.ApplicationID, "amzn1.ask.skill.2604e1db-063f-4f29-a411-6eba27b9dc34"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = system.User.UserID, "amzn1.ask.account.userid"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = system.Device.DeviceID, "amzn1.ask.device.deviceid"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = system.APIEndpoint, "https://api.amazonalexa.com"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = system.APIAccessToken, "reallylongrandomcharacters"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	viewport := alexaRequest.Context.Viewport
	experience := viewport.Experiences[0]

	actual, expected = experience.ArcMinuteWidth, 246
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = experience.ArcMinuteHeight, 144
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = experience.CanRotate, false
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = experience.CanResize, false
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = experience.CanResize, false
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = viewport.Shape, Rectangle.String()
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = viewport.Theme, Light.String()
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = viewport.PixelWidth, 1024
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = viewport.PixelHeight, 600
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = viewport.DPI, 160
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = viewport.CurrentPixelWidth, 1024
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = viewport.CurrentPixelHeight, 600
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = viewport.Touch[0], Single.String()
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = viewport.Keyboard[0], Direction.String()
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	request := alexaRequest.Request
	actual, expected = request.Type, IntentRequest.String()
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = request.RequestID, "amzn1.echo-api.request.d60f9912-bf40-48ab-aa53-c38f46e6fb1f"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	tt, _ := time.Parse(time.RFC3339, "2019-02-23T05:26:19Z")
	actual, expected = request.Timestamp, tt
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = request.Locale, EnUs.String()
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	intent := request.Intent
	actual, expected = intent.Name, "MyCustomIntent"
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}

	actual, expected = intent.ConfirmationStatus, None.String()
	if actual != expected {
		t.Errorf("'%s' != '%s'", actual, expected)
	}
}
