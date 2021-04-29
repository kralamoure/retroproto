// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type HousesRequestLeave struct{}

func (m HousesRequestLeave) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.HousesRequestLeave
}

func (m HousesRequestLeave) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *HousesRequestLeave) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
