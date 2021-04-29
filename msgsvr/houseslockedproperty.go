// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type HousesLockedProperty struct{}

func (m HousesLockedProperty) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.HousesLockedProperty
}

func (m HousesLockedProperty) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *HousesLockedProperty) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
