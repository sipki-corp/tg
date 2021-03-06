package domain

// InputMediaAnimation Represents an animation file (GIF or H.
// 264/MPEG-4 AVC video without sound) to be sent.
type InputMediaAnimation struct {
	// Type of the result, must be animation.
	Type string `json:"type"`
	// Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side.
	// The thumbnail should be in JPEG format and less than 200 kB in size.
	// A thumbnail's width and height should not exceed 320.
	// Ignored if the file is not uploaded using multipart/form-data.
	// Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>.
	// More info on Sending Files ».
	//
	// Optional.
	Thumb string `json:"thumb"`
	// Animation duration.
	//
	// Optional.
	Duration int `json:"duration"`
	// File to send.
	// Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name.
	// More info on Sending Files ».
	Media string `json:"media"`
	// Caption of the animation to be sent, 0-1024 characters after entities parsing.
	//
	// Optional.
	Caption string `json:"caption"`
	// Mode for parsing entities in the animation caption.
	// See formatting options for more details.
	//
	// Optional.
	ParseMode string `json:"parse_mode"`
	// List of special entities that appear in the caption, which can be specified instead of parse_mode.
	//
	// Optional.
	CaptionEntities []MessageEntity `json:"caption_entities"`
	// Animation width.
	//
	// Optional.
	Width int `json:"width"`
	// Animation height.
	//
	// Optional.
	Height int `json:"height"`
}
