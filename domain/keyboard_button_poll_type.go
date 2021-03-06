package domain

// KeyboardButtonPollType This object represents type of poll, which is allowed to be created and sent when the corresponding button is pressed.
type KeyboardButtonPollType struct {
	// If quiz is passed, the user will be allowed to create only polls in the quiz mode.
	// If regular is passed, only regular polls will be allowed.
	// Otherwise, the user will be allowed to create a poll of any type.
	//
	// Optional.
	Type string `json:"type"`
}
