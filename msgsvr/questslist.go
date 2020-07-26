// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type QuestsList struct{}

func (m QuestsList) ProtocolId() d1proto.MsgSvrId {
	return d1proto.QuestsList
}

func (m QuestsList) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *QuestsList) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
