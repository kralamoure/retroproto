// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type QuestGetStep struct{}

func (m QuestGetStep) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.QuestGetStep
}

func (m QuestGetStep) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *QuestGetStep) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
