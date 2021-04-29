// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type QuestGetList struct{}

func (m QuestGetList) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.QuestGetList
}

func (m QuestGetList) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *QuestGetList) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
