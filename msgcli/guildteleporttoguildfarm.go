// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type GuildTeleportToGuildFarm struct{}

func (m GuildTeleportToGuildFarm) ProtocolId() retroproto.MsgCliId {
	return retroproto.GuildTeleportToGuildFarm
}

func (m GuildTeleportToGuildFarm) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildTeleportToGuildFarm) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
