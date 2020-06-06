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
	return "", nil
}

func (m *GameCellData) Deserialize(extra string) error {
	return nil
}
