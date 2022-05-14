// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type GameGetMapData struct{}

func (m GameGetMapData) MessageId() retroproto.MsgCliId {
	return retroproto.GameGetMapData
}

func (m GameGetMapData) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameGetMapData) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
