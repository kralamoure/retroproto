// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type GuildTeleportToGuildHouse struct{}

func (m GuildTeleportToGuildHouse) ProtocolId() retroproto.MsgCliId {
	return retroproto.GuildTeleportToGuildHouse
}

func (m GuildTeleportToGuildHouse) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildTeleportToGuildHouse) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
