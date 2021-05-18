// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GameEffect struct{}

func (m GameEffect) ProtocolId() retroproto.MsgSvrId {
	return retroproto.GameEffect
}

func (m GameEffect) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameEffect) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
