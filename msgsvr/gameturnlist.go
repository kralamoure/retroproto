// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GameTurnList struct{}

func (m GameTurnList) MessageId() retroproto.MsgSvrId {
	return retroproto.GameTurnList
}

func (m GameTurnList) MessageName() string {
	return "GameTurnList"
}

func (m GameTurnList) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameTurnList) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
