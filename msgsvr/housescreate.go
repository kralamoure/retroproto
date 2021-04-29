// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type HousesCreate struct{}

func (m HousesCreate) ProtocolId() d1proto.MsgSvrId {
	return d1proto.HousesCreate
}

func (m HousesCreate) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *HousesCreate) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
