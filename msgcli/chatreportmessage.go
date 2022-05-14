// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type ChatReportMessage struct{}

func (m ChatReportMessage) MessageId() retroproto.MsgCliId {
	return retroproto.ChatReportMessage
}

func (m ChatReportMessage) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ChatReportMessage) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
