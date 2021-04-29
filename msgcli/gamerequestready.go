// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type GameRequestReady struct{}

func (m GameRequestReady) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.GameRequestReady
}

func (m GameRequestReady) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GameRequestReady) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
