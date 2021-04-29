// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type GuildAcceptInvitation struct{}

func (m GuildAcceptInvitation) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.GuildAcceptInvitation
}

func (m GuildAcceptInvitation) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GuildAcceptInvitation) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
