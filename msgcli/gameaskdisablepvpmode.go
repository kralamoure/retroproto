// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type GameAskDisablePVPMode struct{}

func (m GameAskDisablePVPMode) ProtocolId() d1proto.MsgCliId {
	return d1proto.GameAskDisablePVPMode
}

func (m GameAskDisablePVPMode) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *GameAskDisablePVPMode) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
