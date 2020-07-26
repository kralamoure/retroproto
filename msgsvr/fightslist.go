// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type FightsList struct{}

func (m FightsList) ProtocolId() d1proto.MsgSvrId {
	return d1proto.FightsList
}

func (m FightsList) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *FightsList) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
