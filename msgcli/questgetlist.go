// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type QuestGetList struct{}

func (m QuestGetList) ProtocolId() retroproto.MsgCliId {
	return retroproto.QuestGetList
}

func (m QuestGetList) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *QuestGetList) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
