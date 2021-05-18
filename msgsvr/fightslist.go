// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type FightsList struct{}

func (m FightsList) ProtocolId() retroproto.MsgSvrId {
	return retroproto.FightsList
}

func (m FightsList) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *FightsList) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
