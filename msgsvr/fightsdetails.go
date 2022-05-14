// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type FightsDetails struct{}

func NewFightsDetails(extra string) (FightsDetails, error) {
	var m FightsDetails

	if err := m.Deserialize(extra); err != nil {
		return FightsDetails{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m FightsDetails) MessageId() retroproto.MsgSvrId {
	return retroproto.FightsDetails
}

func (m FightsDetails) MessageName() string {
	return "FightsDetails"
}

func (m FightsDetails) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *FightsDetails) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
