// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GuildInvite struct{}

func NewGuildInvite(extra string) (GuildInvite, error) {
	var m GuildInvite

	if err := m.Deserialize(extra); err != nil {
		return GuildInvite{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GuildInvite) MessageId() retroproto.MsgCliId {
	return retroproto.GuildInvite
}

func (m GuildInvite) MessageName() string {
	return "GuildInvite"
}

func (m GuildInvite) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildInvite) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
