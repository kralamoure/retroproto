// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GameZoneData struct{}

func (m GameZoneData) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GameZoneData
}

func (m GameZoneData) Serialized() (string, error) {
	return "", nil
}

func (m *GameZoneData) Deserialize(extra string) error {
	return nil
}
