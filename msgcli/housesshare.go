// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type HousesShare struct{}

func NewHousesShare(extra string) (HousesShare, error) {
	var m HousesShare

	if err := m.Deserialize(extra); err != nil {
		return HousesShare{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m HousesShare) MessageId() retroproto.MsgCliId {
	return retroproto.HousesShare
}

func (m HousesShare) MessageName() string {
	return "HousesShare"
}

func (m HousesShare) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *HousesShare) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
