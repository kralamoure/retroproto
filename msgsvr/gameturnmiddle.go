// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GameTurnMiddle struct{}

func (m GameTurnMiddle) MessageId() retroproto.MsgSvrId {
	return retroproto.GameTurnMiddle
}

func (m GameTurnMiddle) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameTurnMiddle) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
