// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type PartyFollowAll struct{}

func NewPartyFollowAll(extra string) (PartyFollowAll, error) {
	var m PartyFollowAll

	if err := m.Deserialize(extra); err != nil {
		return PartyFollowAll{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m PartyFollowAll) MessageId() retroproto.MsgCliId {
	return retroproto.PartyFollowAll
}

func (m PartyFollowAll) MessageName() string {
	return "PartyFollowAll"
}

func (m PartyFollowAll) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *PartyFollowAll) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
