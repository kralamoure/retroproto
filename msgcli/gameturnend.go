// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type GameTurnEnd struct{}

func (m GameTurnEnd) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.GameTurnEnd
}

func (m GameTurnEnd) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GameTurnEnd) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
