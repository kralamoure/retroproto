// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type FightsBlockJoinerExceptParty struct{}

func NewFightsBlockJoinerExceptParty(extra string) (FightsBlockJoinerExceptParty, error) {
	var m FightsBlockJoinerExceptParty

	if err := m.Deserialize(extra); err != nil {
		return FightsBlockJoinerExceptParty{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m FightsBlockJoinerExceptParty) MessageId() retroproto.MsgCliId {
	return retroproto.FightsBlockJoinerExceptParty
}

func (m FightsBlockJoinerExceptParty) MessageName() string {
	return "FightsBlockJoinerExceptParty"
}

func (m FightsBlockJoinerExceptParty) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *FightsBlockJoinerExceptParty) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
