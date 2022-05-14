// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type GuildGetInfosGuildHouses struct{}

func (m GuildGetInfosGuildHouses) MessageId() retroproto.MsgCliId {
	return retroproto.GuildGetInfosGuildHouses
}

func (m GuildGetInfosGuildHouses) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildGetInfosGuildHouses) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
