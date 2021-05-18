// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type HousesShare struct{}

func (m HousesShare) ProtocolId() retroproto.MsgCliId {
	return retroproto.HousesShare
}

func (m HousesShare) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *HousesShare) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
