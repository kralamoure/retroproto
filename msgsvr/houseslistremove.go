// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type HousesListRemove struct{}

func (m HousesListRemove) ProtocolId() d1proto.MsgSvrId {
	return d1proto.HousesListRemove
}

func (m HousesListRemove) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *HousesListRemove) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
