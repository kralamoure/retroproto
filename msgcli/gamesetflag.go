// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type GameSetFlag struct{}

func (m GameSetFlag) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.GameSetFlag
}

func (m GameSetFlag) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GameSetFlag) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
