package rpc

import (
	"errors"
	"regexp"
	"strconv"
)

func ParseMessage(msg string) (messageType int, err error) {
	messageTypeRegExp := regexp.MustCompile(`^\[(?P<messageId>\d+),(.*)\]$`)
	match := messageTypeRegExp.FindStringSubmatch(msg)

	// Handle malformatted calls
	if len(match) == 0 {
		err = errors.New("malformatted message")
		return
	}

	// Get the messageId
	result := make(map[string]string)
	for i, name := range messageTypeRegExp.SubexpNames() {
		result[name] = match[i]
	}

	messageType, err = strconv.Atoi(result["messageId"])
	if err != nil {
		return
	}

	if messageType < 2 {
		messageType = 0
		err = errors.New("unknown message type")
		return
	} else if messageType > 4 {
		messageType = 0
		err = errors.New("unknown message type")
		return
	}

	return
}
