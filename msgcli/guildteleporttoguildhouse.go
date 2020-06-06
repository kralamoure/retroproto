// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type GuildTeleportToGuildHouse struct{}

func (m GuildTeleportToGuildHouse) ProtocolId() d1proto.MsgCliId {
	return d1proto.GuildTeleportToGuildHouse
}

func (m GuildTeleportToGuildHouse) Serialized() (string, error) {
	return "", nil
}

func (m *GuildTeleportToGuildHouse) Deserialize(extra string) error {
	return nil
}
