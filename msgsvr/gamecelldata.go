// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GameCellData struct{}

func (m GameCellData) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GameCellData
}

func (m GameCellData) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GameCellData) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
