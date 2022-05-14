// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GuildAcceptInvitation struct{}

func NewGuildAcceptInvitation(extra string) (GuildAcceptInvitation, error) {
	var m GuildAcceptInvitation

	if err := m.Deserialize(extra); err != nil {
		return GuildAcceptInvitation{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GuildAcceptInvitation) MessageId() retroproto.MsgCliId {
	return retroproto.GuildAcceptInvitation
}

func (m GuildAcceptInvitation) MessageName() string {
	return "GuildAcceptInvitation"
}

func (m GuildAcceptInvitation) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildAcceptInvitation) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
