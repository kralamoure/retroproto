// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GameTurnList struct{}

func (m GameTurnList) ProtocolId() retroproto.MsgSvrId {
	return retroproto.GameTurnList
}

func (m GameTurnList) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameTurnList) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
