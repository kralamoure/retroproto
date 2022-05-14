// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type GameSetPlayerPosition struct{}

func (m GameSetPlayerPosition) MessageId() retroproto.MsgCliId {
	return retroproto.GameSetPlayerPosition
}

func (m GameSetPlayerPosition) MessageName() string {
	return "GameSetPlayerPosition"
}

func (m GameSetPlayerPosition) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameSetPlayerPosition) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
