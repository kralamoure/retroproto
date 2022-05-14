// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type HousesSell struct{}

func (m HousesSell) MessageId() retroproto.MsgCliId {
	return retroproto.HousesSell
}

func (m HousesSell) MessageName() string {
	return "HousesSell"
}

func (m HousesSell) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *HousesSell) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
