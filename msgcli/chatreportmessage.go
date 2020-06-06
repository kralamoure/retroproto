// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ChatReportMessage struct{}

func (m ChatReportMessage) ProtocolId() d1proto.MsgCliId {
	return d1proto.ChatReportMessage
}

func (m ChatReportMessage) Serialized() (string, error) {
	return "", nil
}

func (m *ChatReportMessage) Deserialize(extra string) error {
	return nil
}
