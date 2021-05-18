// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type HousesRequestLeave struct{}

func (m HousesRequestLeave) ProtocolId() retroproto.MsgCliId {
	return retroproto.HousesRequestLeave
}

func (m HousesRequestLeave) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *HousesRequestLeave) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
