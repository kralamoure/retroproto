// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GameStartToPlay struct{}

func (m GameStartToPlay) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GameStartToPlay
}

func (m GameStartToPlay) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GameStartToPlay) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
