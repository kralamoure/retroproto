// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type GameSetFlag struct{}

func (m GameSetFlag) MessageId() retroproto.MsgCliId {
	return retroproto.GameSetFlag
}

func (m GameSetFlag) MessageName() string {
	return "GameSetFlag"
}

func (m GameSetFlag) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameSetFlag) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
