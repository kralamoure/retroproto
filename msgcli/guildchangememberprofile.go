// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GuildChangeMemberProfile struct{}

func NewGuildChangeMemberProfile(extra string) (GuildChangeMemberProfile, error) {
	var m GuildChangeMemberProfile

	if err := m.Deserialize(extra); err != nil {
		return GuildChangeMemberProfile{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GuildChangeMemberProfile) MessageId() retroproto.MsgCliId {
	return retroproto.GuildChangeMemberProfile
}

func (m GuildChangeMemberProfile) MessageName() string {
	return "GuildChangeMemberProfile"
}

func (m GuildChangeMemberProfile) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildChangeMemberProfile) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
