// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type HousesSell struct{}

func NewHousesSell(extra string) (HousesSell, error) {
	var m HousesSell

	if err := m.Deserialize(extra); err != nil {
		return HousesSell{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

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
