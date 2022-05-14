// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type PartyMovement struct{}

func NewPartyMovement(extra string) (PartyMovement, error) {
	var m PartyMovement

	if err := m.Deserialize(extra); err != nil {
		return PartyMovement{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m PartyMovement) MessageId() retroproto.MsgSvrId {
	return retroproto.PartyMovement
}

func (m PartyMovement) MessageName() string {
	return "PartyMovement"
}

func (m PartyMovement) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *PartyMovement) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
