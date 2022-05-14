// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type QuestsList struct{}

func (m QuestsList) MessageId() retroproto.MsgSvrId {
	return retroproto.QuestsList
}

func (m QuestsList) MessageName() string {
	return "QuestsList"
}

func (m QuestsList) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *QuestsList) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
