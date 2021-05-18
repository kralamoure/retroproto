// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type GameSetPlayerPosition struct{}

func (m GameSetPlayerPosition) ProtocolId() retroproto.MsgCliId {
	return retroproto.GameSetPlayerPosition
}

func (m GameSetPlayerPosition) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameSetPlayerPosition) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
