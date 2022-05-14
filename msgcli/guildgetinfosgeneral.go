// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type GuildGetInfosGeneral struct{}

func (m GuildGetInfosGeneral) MessageId() retroproto.MsgCliId {
	return retroproto.GuildGetInfosGeneral
}

func (m GuildGetInfosGeneral) MessageName() string {
	return "GuildGetInfosGeneral"
}

func (m GuildGetInfosGeneral) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildGetInfosGeneral) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
