package domain

// ResponseParameters Contains information about why a request was unsuccessful.
type ResponseParameters struct {
	// The group has been migrated to a supergroup with the specified identifier.
	// This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it.
	// But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
	//
	// Optional.
	MigrateToChatId int `json:"migrate_to_chat_id"`
	// In case of exceeding flood control, the number of seconds left to wait before the request can be repeated.
	//
	// Optional.
	RetryAfter int `json:"retry_after"`
}
