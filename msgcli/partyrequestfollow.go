// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type PartyRequestFollow struct{}

func NewPartyRequestFollow(extra string) (PartyRequestFollow, error) {
	var m PartyRequestFollow

	if err := m.Deserialize(extra); err != nil {
		return PartyRequestFollow{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m PartyRequestFollow) MessageId() retroproto.MsgCliId {
	return retroproto.PartyRequestFollow
}

func (m PartyRequestFollow) MessageName() string {
	return "PartyRequestFollow"
}

func (m PartyRequestFollow) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *PartyRequestFollow) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
