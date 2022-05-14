// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type PartyAcceptInvitation struct{}

func NewPartyAcceptInvitation(extra string) (PartyAcceptInvitation, error) {
	var m PartyAcceptInvitation

	if err := m.Deserialize(extra); err != nil {
		return PartyAcceptInvitation{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m PartyAcceptInvitation) MessageId() retroproto.MsgCliId {
	return retroproto.PartyAcceptInvitation
}

func (m PartyAcceptInvitation) MessageName() string {
	return "PartyAcceptInvitation"
}

func (m PartyAcceptInvitation) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *PartyAcceptInvitation) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
