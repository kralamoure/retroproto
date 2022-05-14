// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type HousesCreate struct{}

func NewHousesCreate(extra string) (HousesCreate, error) {
	var m HousesCreate

	if err := m.Deserialize(extra); err != nil {
		return HousesCreate{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m HousesCreate) MessageId() retroproto.MsgSvrId {
	return retroproto.HousesCreate
}

func (m HousesCreate) MessageName() string {
	return "HousesCreate"
}

func (m HousesCreate) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *HousesCreate) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
