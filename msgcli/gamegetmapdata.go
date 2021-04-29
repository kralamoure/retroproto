// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type GameGetMapData struct{}

func (m GameGetMapData) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.GameGetMapData
}

func (m GameGetMapData) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GameGetMapData) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
