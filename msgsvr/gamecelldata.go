// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GameCellData struct{}

func (m GameCellData) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GameCellData
}

func (m GameCellData) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *GameCellData) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
