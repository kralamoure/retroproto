// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type PartyRefuse struct{}

func NewPartyRefuse(extra string) (PartyRefuse, error) {
	var m PartyRefuse

	if err := m.Deserialize(extra); err != nil {
		return PartyRefuse{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m PartyRefuse) MessageId() retroproto.MsgSvrId {
	return retroproto.PartyRefuse
}

func (m PartyRefuse) MessageName() string {
	return "PartyRefuse"
}

func (m PartyRefuse) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *PartyRefuse) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
