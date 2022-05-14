// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type GuildRefuseInvitation struct{}

func (m GuildRefuseInvitation) MessageId() retroproto.MsgCliId {
	return retroproto.GuildRefuseInvitation
}

func (m GuildRefuseInvitation) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildRefuseInvitation) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
