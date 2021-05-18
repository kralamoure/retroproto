// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type HousesProperties struct{}

func (m HousesProperties) ProtocolId() retroproto.MsgSvrId {
	return retroproto.HousesProperties
}

func (m HousesProperties) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *HousesProperties) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
