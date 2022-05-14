// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GuildInfosGeneral struct{}

func (m GuildInfosGeneral) MessageId() retroproto.MsgSvrId {
	return retroproto.GuildInfosGeneral
}

func (m GuildInfosGeneral) MessageName() string {
	return "GuildInfosGeneral"
}

func (m GuildInfosGeneral) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildInfosGeneral) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
