// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GamePositionStart struct{}

func (m GamePositionStart) MessageId() retroproto.MsgSvrId {
	return retroproto.GamePositionStart
}

func (m GamePositionStart) MessageName() string {
	return "GamePositionStart"
}

func (m GamePositionStart) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GamePositionStart) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
