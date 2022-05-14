// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type PartyRequestLeave struct{}

func NewPartyRequestLeave(extra string) (PartyRequestLeave, error) {
	var m PartyRequestLeave

	if err := m.Deserialize(extra); err != nil {
		return PartyRequestLeave{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m PartyRequestLeave) MessageId() retroproto.MsgCliId {
	return retroproto.PartyRequestLeave
}

func (m PartyRequestLeave) MessageName() string {
	return "PartyRequestLeave"
}

func (m PartyRequestLeave) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *PartyRequestLeave) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
