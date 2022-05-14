// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GameGameOver struct{}

func (m GameGameOver) MessageId() retroproto.MsgSvrId {
	return retroproto.GameGameOver
}

func (m GameGameOver) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameGameOver) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
