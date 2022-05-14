// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type PartyLeave struct{}

func NewPartyLeave(extra string) (PartyLeave, error) {
	var m PartyLeave

	if err := m.Deserialize(extra); err != nil {
		return PartyLeave{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m PartyLeave) MessageId() retroproto.MsgSvrId {
	return retroproto.PartyLeave
}

func (m PartyLeave) MessageName() string {
	return "PartyLeave"
}

func (m PartyLeave) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *PartyLeave) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
