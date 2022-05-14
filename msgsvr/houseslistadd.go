// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type HousesListAdd struct{}

func NewHousesListAdd(extra string) (HousesListAdd, error) {
	var m HousesListAdd

	if err := m.Deserialize(extra); err != nil {
		return HousesListAdd{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m HousesListAdd) MessageId() retroproto.MsgSvrId {
	return retroproto.HousesListAdd
}

func (m HousesListAdd) MessageName() string {
	return "HousesListAdd"
}

func (m HousesListAdd) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *HousesListAdd) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
