// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type HousesProperties struct{}

func NewHousesProperties(extra string) (HousesProperties, error) {
	var m HousesProperties

	if err := m.Deserialize(extra); err != nil {
		return HousesProperties{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m HousesProperties) MessageId() retroproto.MsgSvrId {
	return retroproto.HousesProperties
}

func (m HousesProperties) MessageName() string {
	return "HousesProperties"
}

func (m HousesProperties) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *HousesProperties) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
