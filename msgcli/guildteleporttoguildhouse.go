// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GuildTeleportToGuildHouse struct{}

func NewGuildTeleportToGuildHouse(extra string) (GuildTeleportToGuildHouse, error) {
	var m GuildTeleportToGuildHouse

	if err := m.Deserialize(extra); err != nil {
		return GuildTeleportToGuildHouse{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GuildTeleportToGuildHouse) MessageId() retroproto.MsgCliId {
	return retroproto.GuildTeleportToGuildHouse
}

func (m GuildTeleportToGuildHouse) MessageName() string {
	return "GuildTeleportToGuildHouse"
}

func (m GuildTeleportToGuildHouse) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildTeleportToGuildHouse) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
