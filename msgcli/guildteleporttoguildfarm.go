// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type GuildTeleportToGuildFarm struct{}

func (m GuildTeleportToGuildFarm) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.GuildTeleportToGuildFarm
}

func (m GuildTeleportToGuildFarm) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GuildTeleportToGuildFarm) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
