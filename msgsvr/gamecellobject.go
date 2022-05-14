// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GameCellObject struct{}

func (m GameCellObject) MessageId() retroproto.MsgSvrId {
	return retroproto.GameCellObject
}

func (m GameCellObject) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameCellObject) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
