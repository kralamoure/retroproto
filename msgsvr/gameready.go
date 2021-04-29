// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GameReady struct{}

func (m GameReady) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GameReady
}

func (m GameReady) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GameReady) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
