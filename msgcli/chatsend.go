package msgcli

import (
	"fmt"
	"strings"

	"github.com/kralamoure/d1proto"
)

type ChatSend struct {
	Destination string
	Message     string
}

func (m ChatSend) ProtocolId() d1proto.MsgCliId {
	return d1proto.ChatSend
}

func (m ChatSend) Serialized() (string, error) {
	return fmt.Sprintf("%s|%s", m.Destination, m.Message), nil
}

func (m *ChatSend) Deserialize(extra string) error {
	sli := strings.SplitN(extra, "|", 2)
	if len(sli) != 2 {
		return d1proto.ErrInvalidMsg
	}

	m.Destination = sli[0]
	m.Message = sli[1]

	return nil
}
