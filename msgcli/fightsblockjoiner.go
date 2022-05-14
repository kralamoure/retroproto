// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type FightsBlockJoiner struct{}

func NewFightsBlockJoiner(extra string) (FightsBlockJoiner, error) {
	var m FightsBlockJoiner

	if err := m.Deserialize(extra); err != nil {
		return FightsBlockJoiner{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m FightsBlockJoiner) MessageId() retroproto.MsgCliId {
	return retroproto.FightsBlockJoiner
}

func (m FightsBlockJoiner) MessageName() string {
	return "FightsBlockJoiner"
}

func (m FightsBlockJoiner) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *FightsBlockJoiner) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
