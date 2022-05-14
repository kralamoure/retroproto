// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type FightsGetDetails struct{}

func NewFightsGetDetails(extra string) (FightsGetDetails, error) {
	var m FightsGetDetails

	if err := m.Deserialize(extra); err != nil {
		return FightsGetDetails{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m FightsGetDetails) MessageId() retroproto.MsgCliId {
	return retroproto.FightsGetDetails
}

func (m FightsGetDetails) MessageName() string {
	return "FightsGetDetails"
}

func (m FightsGetDetails) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *FightsGetDetails) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
