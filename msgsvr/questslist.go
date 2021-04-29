// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type QuestsList struct{}

func (m QuestsList) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.QuestsList
}

func (m QuestsList) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *QuestsList) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
