// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type FightsBlockSpectators struct{}

func NewFightsBlockSpectators(extra string) (FightsBlockSpectators, error) {
	var m FightsBlockSpectators

	if err := m.Deserialize(extra); err != nil {
		return FightsBlockSpectators{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m FightsBlockSpectators) MessageId() retroproto.MsgCliId {
	return retroproto.FightsBlockSpectators
}

func (m FightsBlockSpectators) MessageName() string {
	return "FightsBlockSpectators"
}

func (m FightsBlockSpectators) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *FightsBlockSpectators) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
