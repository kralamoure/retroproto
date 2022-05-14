// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GameCellData struct{}

func (m GameCellData) MessageId() retroproto.MsgSvrId {
	return retroproto.GameCellData
}

func (m GameCellData) MessageName() string {
	return "GameCellData"
}

func (m GameCellData) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameCellData) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
