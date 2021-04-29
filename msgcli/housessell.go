// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type HousesSell struct{}

func (m HousesSell) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.HousesSell
}

func (m HousesSell) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *HousesSell) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
