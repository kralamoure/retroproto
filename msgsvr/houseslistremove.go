// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type HousesListRemove struct{}

func NewHousesListRemove(extra string) (HousesListRemove, error) {
	var m HousesListRemove

	if err := m.Deserialize(extra); err != nil {
		return HousesListRemove{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m HousesListRemove) MessageId() retroproto.MsgSvrId {
	return retroproto.HousesListRemove
}

func (m HousesListRemove) MessageName() string {
	return "HousesListRemove"
}

func (m HousesListRemove) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *HousesListRemove) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
