// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type PartyRefuseInvitation struct{}

func NewPartyRefuseInvitation(extra string) (PartyRefuseInvitation, error) {
	var m PartyRefuseInvitation

	if err := m.Deserialize(extra); err != nil {
		return PartyRefuseInvitation{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m PartyRefuseInvitation) MessageId() retroproto.MsgCliId {
	return retroproto.PartyRefuseInvitation
}

func (m PartyRefuseInvitation) MessageName() string {
	return "PartyRefuseInvitation"
}

func (m PartyRefuseInvitation) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *PartyRefuseInvitation) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
