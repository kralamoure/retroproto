// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GamePositionStart struct{}

func (m GamePositionStart) ProtocolId() retroproto.MsgSvrId {
	return retroproto.GamePositionStart
}

func (m GamePositionStart) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GamePositionStart) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
