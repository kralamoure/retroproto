// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type HousesState struct{}

func (m HousesState) ProtocolId() d1proto.MsgCliId {
	return d1proto.HousesState
}

func (m HousesState) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *HousesState) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}