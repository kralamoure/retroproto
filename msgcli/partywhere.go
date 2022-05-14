// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type PartyWhere struct{}

func NewPartyWhere(extra string) (PartyWhere, error) {
	var m PartyWhere

	if err := m.Deserialize(extra); err != nil {
		return PartyWhere{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m PartyWhere) MessageId() retroproto.MsgCliId {
	return retroproto.PartyWhere
}

func (m PartyWhere) MessageName() string {
	return "PartyWhere"
}

func (m PartyWhere) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *PartyWhere) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
