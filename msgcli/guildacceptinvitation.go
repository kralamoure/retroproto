// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type GuildAcceptInvitation struct{}

func (m GuildAcceptInvitation) ProtocolId() d1proto.MsgCliId {
	return d1proto.GuildAcceptInvitation
}

func (m GuildAcceptInvitation) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *GuildAcceptInvitation) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
