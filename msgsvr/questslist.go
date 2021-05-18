// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type QuestsList struct{}

func (m QuestsList) ProtocolId() retroproto.MsgSvrId {
	return retroproto.QuestsList
}

func (m QuestsList) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *QuestsList) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
