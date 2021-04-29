// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type HousesListAdd struct{}

func (m HousesListAdd) ProtocolId() d1proto.MsgSvrId {
	return d1proto.HousesListAdd
}

func (m HousesListAdd) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *HousesListAdd) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
