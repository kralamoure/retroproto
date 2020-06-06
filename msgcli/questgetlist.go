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
	return "", nil
}

func (m *QuestGetList) Deserialize(extra string) error {
	return nil
}
