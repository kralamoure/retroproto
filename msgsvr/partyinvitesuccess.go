// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type PartyInviteSuccess struct{}

func NewPartyInviteSuccess(extra string) (PartyInviteSuccess, error) {
	var m PartyInviteSuccess

	if err := m.Deserialize(extra); err != nil {
		return PartyInviteSuccess{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m PartyInviteSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.PartyInviteSuccess
}

func (m PartyInviteSuccess) MessageName() string {
	return "PartyInviteSuccess"
}

func (m PartyInviteSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *PartyInviteSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
