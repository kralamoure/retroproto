// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type GuildTeleportToGuildFarm struct{}

func (m GuildTeleportToGuildFarm) ProtocolId() d1proto.MsgCliId {
	return d1proto.GuildTeleportToGuildFarm
}

func (m GuildTeleportToGuildFarm) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *GuildTeleportToGuildFarm) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}