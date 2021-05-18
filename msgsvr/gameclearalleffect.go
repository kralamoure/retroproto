// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GameClearAllEffect struct{}

func (m GameClearAllEffect) ProtocolId() retroproto.MsgSvrId {
	return retroproto.GameClearAllEffect
}

func (m GameClearAllEffect) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameClearAllEffect) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
