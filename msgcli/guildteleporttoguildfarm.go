// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GuildTeleportToGuildFarm struct{}

func NewGuildTeleportToGuildFarm(extra string) (GuildTeleportToGuildFarm, error) {
	var m GuildTeleportToGuildFarm

	if err := m.Deserialize(extra); err != nil {
		return GuildTeleportToGuildFarm{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GuildTeleportToGuildFarm) MessageId() retroproto.MsgCliId {
	return retroproto.GuildTeleportToGuildFarm
}

func (m GuildTeleportToGuildFarm) MessageName() string {
	return "GuildTeleportToGuildFarm"
}

func (m GuildTeleportToGuildFarm) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildTeleportToGuildFarm) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
