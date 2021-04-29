// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GameTurnList struct{}

func (m GameTurnList) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GameTurnList
}

func (m GameTurnList) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GameTurnList) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
