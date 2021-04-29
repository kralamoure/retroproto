// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type FightsList struct{}

func (m FightsList) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.FightsList
}

func (m FightsList) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *FightsList) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
