// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type GuildRefuseInvitation struct{}

func (m GuildRefuseInvitation) ProtocolId() d1proto.MsgCliId {
	return d1proto.GuildRefuseInvitation
}

func (m GuildRefuseInvitation) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *GuildRefuseInvitation) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
