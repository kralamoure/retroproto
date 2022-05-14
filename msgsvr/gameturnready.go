// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GameTurnReady struct{}

func (m GameTurnReady) MessageId() retroproto.MsgSvrId {
	return retroproto.GameTurnReady
}

func (m GameTurnReady) MessageName() string {
	return "GameTurnReady"
}

func (m GameTurnReady) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameTurnReady) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
