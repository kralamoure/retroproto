// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type GameSetPlayerPosition struct{}

func (m GameSetPlayerPosition) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.GameSetPlayerPosition
}

func (m GameSetPlayerPosition) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GameSetPlayerPosition) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
