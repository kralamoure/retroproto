// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type GameAskDisablePVPMode struct{}

func (m GameAskDisablePVPMode) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.GameAskDisablePVPMode
}

func (m GameAskDisablePVPMode) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GameAskDisablePVPMode) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
