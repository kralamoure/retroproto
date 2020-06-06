// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type QuestGetStep struct{}

func (m QuestGetStep) ProtocolId() d1proto.MsgCliId {
	return d1proto.QuestGetStep
}

func (m QuestGetStep) Serialized() (string, error) {
	return "", nil
}

func (m *QuestGetStep) Deserialize(extra string) error {
	return nil
}
