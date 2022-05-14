// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GameEffect struct{}

func (m GameEffect) MessageId() retroproto.MsgSvrId {
	return retroproto.GameEffect
}

func (m GameEffect) MessageName() string {
	return "GameEffect"
}

func (m GameEffect) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameEffect) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
