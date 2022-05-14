// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type HousesLeave struct{}

func NewHousesLeave(extra string) (HousesLeave, error) {
	var m HousesLeave

	if err := m.Deserialize(extra); err != nil {
		return HousesLeave{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m HousesLeave) MessageId() retroproto.MsgSvrId {
	return retroproto.HousesLeave
}

func (m HousesLeave) MessageName() string {
	return "HousesLeave"
}

func (m HousesLeave) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *HousesLeave) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
