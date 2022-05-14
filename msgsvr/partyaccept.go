// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type PartyAccept struct{}

func NewPartyAccept(extra string) (PartyAccept, error) {
	var m PartyAccept

	if err := m.Deserialize(extra); err != nil {
		return PartyAccept{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m PartyAccept) MessageId() retroproto.MsgSvrId {
	return retroproto.PartyAccept
}

func (m PartyAccept) MessageName() string {
	return "PartyAccept"
}

func (m PartyAccept) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *PartyAccept) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
