// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type HousesUnShare struct{}

func (m HousesUnShare) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.HousesUnShare
}

func (m HousesUnShare) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *HousesUnShare) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
