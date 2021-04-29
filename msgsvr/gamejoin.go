// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GameJoin struct{}

func (m GameJoin) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GameJoin
}

func (m GameJoin) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *GameJoin) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
