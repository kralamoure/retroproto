// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GameStartToPlay struct{}

func (m GameStartToPlay) MessageId() retroproto.MsgSvrId {
	return retroproto.GameStartToPlay
}

func (m GameStartToPlay) MessageName() string {
	return "GameStartToPlay"
}

func (m GameStartToPlay) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameStartToPlay) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
