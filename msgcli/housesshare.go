// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type HousesShare struct{}

func (m HousesShare) ProtocolId() d1proto.MsgCliId {
	return d1proto.HousesShare
}

func (m HousesShare) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *HousesShare) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
