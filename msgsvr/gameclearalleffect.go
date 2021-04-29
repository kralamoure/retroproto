// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GameClearAllEffect struct{}

func (m GameClearAllEffect) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GameClearAllEffect
}

func (m GameClearAllEffect) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GameClearAllEffect) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
