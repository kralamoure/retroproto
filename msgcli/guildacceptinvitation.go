// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type GuildAcceptInvitation struct{}

func (m GuildAcceptInvitation) ProtocolId() retroproto.MsgCliId {
	return retroproto.GuildAcceptInvitation
}

func (m GuildAcceptInvitation) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildAcceptInvitation) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
