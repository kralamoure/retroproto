// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type PartyInviteError struct{}

func NewPartyInviteError(extra string) (PartyInviteError, error) {
	var m PartyInviteError

	if err := m.Deserialize(extra); err != nil {
		return PartyInviteError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m PartyInviteError) MessageId() retroproto.MsgSvrId {
	return retroproto.PartyInviteError
}

func (m PartyInviteError) MessageName() string {
	return "PartyInviteError"
}

func (m PartyInviteError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *PartyInviteError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
