// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ConquestSwitchPlaces struct{}

func NewConquestSwitchPlaces(extra string) (ConquestSwitchPlaces, error) {
	var m ConquestSwitchPlaces

	if err := m.Deserialize(extra); err != nil {
		return ConquestSwitchPlaces{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ConquestSwitchPlaces) MessageId() retroproto.MsgCliId {
	return retroproto.ConquestSwitchPlaces
}

func (m ConquestSwitchPlaces) MessageName() string {
	return "ConquestSwitchPlaces"
}

func (m ConquestSwitchPlaces) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ConquestSwitchPlaces) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
