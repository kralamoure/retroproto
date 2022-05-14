// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type HousesLockedProperty struct{}

func NewHousesLockedProperty(extra string) (HousesLockedProperty, error) {
	var m HousesLockedProperty

	if err := m.Deserialize(extra); err != nil {
		return HousesLockedProperty{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m HousesLockedProperty) MessageId() retroproto.MsgSvrId {
	return retroproto.HousesLockedProperty
}

func (m HousesLockedProperty) MessageName() string {
	return "HousesLockedProperty"
}

func (m HousesLockedProperty) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *HousesLockedProperty) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
