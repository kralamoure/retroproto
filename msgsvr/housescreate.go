// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type HousesCreate struct{}

func (m HousesCreate) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.HousesCreate
}

func (m HousesCreate) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *HousesCreate) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
