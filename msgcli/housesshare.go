// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type HousesShare struct{}

func (m HousesShare) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.HousesShare
}

func (m HousesShare) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *HousesShare) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
