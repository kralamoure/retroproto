// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type GameTurnEnd struct{}

func (m GameTurnEnd) MessageId() retroproto.MsgCliId {
	return retroproto.GameTurnEnd
}

func (m GameTurnEnd) MessageName() string {
	return "GameTurnEnd"
}

func (m GameTurnEnd) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameTurnEnd) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
