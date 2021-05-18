// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type HousesUnShare struct{}

func (m HousesUnShare) ProtocolId() retroproto.MsgCliId {
	return retroproto.HousesUnShare
}

func (m HousesUnShare) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *HousesUnShare) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
