// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type HousesKick struct{}

func NewHousesKick(extra string) (HousesKick, error) {
	var m HousesKick

	if err := m.Deserialize(extra); err != nil {
		return HousesKick{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m HousesKick) MessageId() retroproto.MsgCliId {
	return retroproto.HousesKick
}

func (m HousesKick) MessageName() string {
	return "HousesKick"
}

func (m HousesKick) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *HousesKick) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
