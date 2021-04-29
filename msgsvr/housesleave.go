// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type HousesLeave struct{}

func (m HousesLeave) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.HousesLeave
}

func (m HousesLeave) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *HousesLeave) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
