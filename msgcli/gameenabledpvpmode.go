// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type GameEnabledPVPMode struct{}

func (m GameEnabledPVPMode) ProtocolId() d1proto.MsgCliId {
	return d1proto.GameEnabledPVPMode
}

func (m GameEnabledPVPMode) Serialized() (string, error) {
	return "", nil
}

func (m *GameEnabledPVPMode) Deserialize(extra string) error {
	return nil
}
