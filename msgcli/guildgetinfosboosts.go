// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type GuildGetInfosBoosts struct{}

func (m GuildGetInfosBoosts) MessageId() retroproto.MsgCliId {
	return retroproto.GuildGetInfosBoosts
}

func (m GuildGetInfosBoosts) MessageName() string {
	return "GuildGetInfosBoosts"
}

func (m GuildGetInfosBoosts) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildGetInfosBoosts) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
