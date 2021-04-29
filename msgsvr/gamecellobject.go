// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GameCellObject struct{}

func (m GameCellObject) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GameCellObject
}

func (m GameCellObject) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GameCellObject) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
