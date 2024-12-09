package utils

import (
	"chat-database/ent/message"
	"fmt"
)

func ParseTypeMessage(s string) (message.TypeMessage, error) {
	switch s {
	case "text":
		return message.TypeMessageText, nil
	case "image":
		return message.TypeMessageImage, nil
	case "audio":
		return message.TypeMessageAudio, nil
	case "video":
		return message.TypeMessageVideo, nil
	case "location":
		return message.TypeMessageLocation, nil
	case "contact":
		return message.TypeMessageContact, nil
	default:
		return "", fmt.Errorf("invalid TypeMessage: %s", s)
	}
}
