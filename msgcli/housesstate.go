// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type HousesState struct{}

func (m HousesState) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.HousesState
}

func (m HousesState) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *HousesState) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
