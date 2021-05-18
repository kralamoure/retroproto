// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type GameAskDisablePVPMode struct{}

func (m GameAskDisablePVPMode) ProtocolId() retroproto.MsgCliId {
	return retroproto.GameAskDisablePVPMode
}

func (m GameAskDisablePVPMode) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameAskDisablePVPMode) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
