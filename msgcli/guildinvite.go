// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type GuildInvite struct{}

func (m GuildInvite) MessageId() retroproto.MsgCliId {
	return retroproto.GuildInvite
}

func (m GuildInvite) MessageName() string {
	return "GuildInvite"
}

func (m GuildInvite) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildInvite) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
