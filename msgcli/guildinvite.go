// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type GuildInvite struct{}

func (m GuildInvite) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.GuildInvite
}

func (m GuildInvite) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GuildInvite) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
