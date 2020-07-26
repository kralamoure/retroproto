// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type QuestGetList struct{}

func (m QuestGetList) ProtocolId() d1proto.MsgCliId {
	return d1proto.QuestGetList
}

func (m QuestGetList) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *QuestGetList) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
