// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type GuildTeleportToGuildHouse struct{}

func (m GuildTeleportToGuildHouse) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.GuildTeleportToGuildHouse
}

func (m GuildTeleportToGuildHouse) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GuildTeleportToGuildHouse) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
