// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type PartyInvite struct{}

func NewPartyInvite(extra string) (PartyInvite, error) {
	var m PartyInvite

	if err := m.Deserialize(extra); err != nil {
		return PartyInvite{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m PartyInvite) MessageId() retroproto.MsgCliId {
	return retroproto.PartyInvite
}

func (m PartyInvite) MessageName() string {
	return "PartyInvite"
}

func (m PartyInvite) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *PartyInvite) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
