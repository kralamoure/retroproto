// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type GuildGetInfosGuildHouses struct{}

func (m GuildGetInfosGuildHouses) ProtocolId() d1proto.MsgCliId {
	return d1proto.GuildGetInfosGuildHouses
}

func (m GuildGetInfosGuildHouses) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *GuildGetInfosGuildHouses) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
