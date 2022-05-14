// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type GuildAcceptInvitation struct{}

func (m GuildAcceptInvitation) MessageId() retroproto.MsgCliId {
	return retroproto.GuildAcceptInvitation
}

func (m GuildAcceptInvitation) MessageName() string {
	return "GuildAcceptInvitation"
}

func (m GuildAcceptInvitation) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildAcceptInvitation) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
