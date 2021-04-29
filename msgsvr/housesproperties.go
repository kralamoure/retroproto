// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type HousesProperties struct{}

func (m HousesProperties) ProtocolId() d1proto.MsgSvrId {
	return d1proto.HousesProperties
}

func (m HousesProperties) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *HousesProperties) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
