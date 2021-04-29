// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type HousesProperties struct{}

func (m HousesProperties) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.HousesProperties
}

func (m HousesProperties) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *HousesProperties) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
