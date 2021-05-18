// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type GameTurnOk struct{}

func (m GameTurnOk) ProtocolId() retroproto.MsgCliId {
	return retroproto.GameTurnOk
}

func (m GameTurnOk) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameTurnOk) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
