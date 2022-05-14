// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GuildRefuseInvitation struct{}

func NewGuildRefuseInvitation(extra string) (GuildRefuseInvitation, error) {
	var m GuildRefuseInvitation

	if err := m.Deserialize(extra); err != nil {
		return GuildRefuseInvitation{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GuildRefuseInvitation) MessageId() retroproto.MsgCliId {
	return retroproto.GuildRefuseInvitation
}

func (m GuildRefuseInvitation) MessageName() string {
	return "GuildRefuseInvitation"
}

func (m GuildRefuseInvitation) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildRefuseInvitation) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
