// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type GameEnabledPVPMode struct{}

func (m GameEnabledPVPMode) ProtocolId() retroproto.MsgCliId {
	return retroproto.GameEnabledPVPMode
}

func (m GameEnabledPVPMode) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameEnabledPVPMode) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
