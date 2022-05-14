// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type PartyFollowError struct{}

func NewPartyFollowError(extra string) (PartyFollowError, error) {
	var m PartyFollowError

	if err := m.Deserialize(extra); err != nil {
		return PartyFollowError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m PartyFollowError) MessageId() retroproto.MsgSvrId {
	return retroproto.PartyFollowError
}

func (m PartyFollowError) MessageName() string {
	return "PartyFollowError"
}

func (m PartyFollowError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *PartyFollowError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
