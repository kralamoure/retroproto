// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type GameEnabledPVPMode struct{}

func (m GameEnabledPVPMode) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.GameEnabledPVPMode
}

func (m GameEnabledPVPMode) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GameEnabledPVPMode) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
