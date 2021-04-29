// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type GameTurnOk struct{}

func (m GameTurnOk) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.GameTurnOk
}

func (m GameTurnOk) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GameTurnOk) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
