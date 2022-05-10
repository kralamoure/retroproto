package retroproto

import (
	"errors"
)

var errInvalidMessage = errors.New("invalid message")

type Message interface {
	MessageId() string
	String() string
}
