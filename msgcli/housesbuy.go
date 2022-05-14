// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type HousesBuy struct{}

func NewHousesBuy(extra string) (HousesBuy, error) {
	var m HousesBuy

	if err := m.Deserialize(extra); err != nil {
		return HousesBuy{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m HousesBuy) MessageId() retroproto.MsgCliId {
	return retroproto.HousesBuy
}

func (m HousesBuy) MessageName() string {
	return "HousesBuy"
}

func (m HousesBuy) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *HousesBuy) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
