// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GameJoin struct{}

func (m GameJoin) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GameJoin
}

func (m GameJoin) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GameJoin) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
