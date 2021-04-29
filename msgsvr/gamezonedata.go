// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GameZoneData struct{}

func (m GameZoneData) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GameZoneData
}

func (m GameZoneData) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GameZoneData) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
