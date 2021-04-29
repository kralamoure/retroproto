// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type GuildRefuseInvitation struct{}

func (m GuildRefuseInvitation) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.GuildRefuseInvitation
}

func (m GuildRefuseInvitation) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GuildRefuseInvitation) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
