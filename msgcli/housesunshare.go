// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type HousesUnShare struct{}

func NewHousesUnShare(extra string) (HousesUnShare, error) {
	var m HousesUnShare

	if err := m.Deserialize(extra); err != nil {
		return HousesUnShare{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m HousesUnShare) MessageId() retroproto.MsgCliId {
	return retroproto.HousesUnShare
}

func (m HousesUnShare) MessageName() string {
	return "HousesUnShare"
}

func (m HousesUnShare) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *HousesUnShare) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
