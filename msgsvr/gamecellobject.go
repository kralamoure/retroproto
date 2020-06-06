// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GameCellObject struct{}

func (m GameCellObject) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GameCellObject
}

func (m GameCellObject) Serialized() (string, error) {
	return "", nil
}

func (m *GameCellObject) Deserialize(extra string) error {
	return nil
}
