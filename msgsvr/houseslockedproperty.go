// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type HousesLockedProperty struct{}

func (m HousesLockedProperty) ProtocolId() retroproto.MsgSvrId {
	return retroproto.HousesLockedProperty
}

func (m HousesLockedProperty) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *HousesLockedProperty) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
