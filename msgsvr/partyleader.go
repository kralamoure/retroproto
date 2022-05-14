// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type PartyLeader struct{}

func NewPartyLeader(extra string) (PartyLeader, error) {
	var m PartyLeader

	if err := m.Deserialize(extra); err != nil {
		return PartyLeader{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m PartyLeader) MessageId() retroproto.MsgSvrId {
	return retroproto.PartyLeader
}

func (m PartyLeader) MessageName() string {
	return "PartyLeader"
}

func (m PartyLeader) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *PartyLeader) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
