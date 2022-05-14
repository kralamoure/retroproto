// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type PartyCreateError struct{}

func NewPartyCreateError(extra string) (PartyCreateError, error) {
	var m PartyCreateError

	if err := m.Deserialize(extra); err != nil {
		return PartyCreateError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m PartyCreateError) MessageId() retroproto.MsgSvrId {
	return retroproto.PartyCreateError
}

func (m PartyCreateError) MessageName() string {
	return "PartyCreateError"
}

func (m PartyCreateError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *PartyCreateError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
