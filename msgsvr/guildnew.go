// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GuildNew struct{}

func (m GuildNew) MessageId() retroproto.MsgSvrId {
	return retroproto.GuildNew
}

func (m GuildNew) MessageName() string {
	return "GuildNew"
}

func (m GuildNew) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildNew) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
