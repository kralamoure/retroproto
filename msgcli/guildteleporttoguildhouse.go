// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type GuildTeleportToGuildHouse struct{}

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
