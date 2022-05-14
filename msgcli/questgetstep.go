// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type QuestGetStep struct{}

func (m QuestGetStep) MessageId() retroproto.MsgCliId {
	return retroproto.QuestGetStep
}

func (m QuestGetStep) MessageName() string {
	return "QuestGetStep"
}

func (m QuestGetStep) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *QuestGetStep) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
