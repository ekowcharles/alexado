// Package alexado provides objects and basic behavior for the sending requests to and processing responses for Alexa
package alexado

type Attributes struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}
