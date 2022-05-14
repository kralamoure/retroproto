// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ChatReportMessage struct{}

func NewChatReportMessage(extra string) (ChatReportMessage, error) {
	var m ChatReportMessage

	if err := m.Deserialize(extra); err != nil {
		return ChatReportMessage{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ChatReportMessage) MessageId() retroproto.MsgCliId {
	return retroproto.ChatReportMessage
}

func (m ChatReportMessage) MessageName() string {
	return "ChatReportMessage"
}

func (m ChatReportMessage) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ChatReportMessage) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
