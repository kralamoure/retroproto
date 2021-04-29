// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type ChatReportMessage struct{}

func (m ChatReportMessage) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.ChatReportMessage
}

func (m ChatReportMessage) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ChatReportMessage) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
