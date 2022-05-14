// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GameJoin struct{}

func (m GameJoin) MessageId() retroproto.MsgSvrId {
	return retroproto.GameJoin
}

func (m GameJoin) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameJoin) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
