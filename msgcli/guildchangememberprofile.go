// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type GuildChangeMemberProfile struct{}

func (m GuildChangeMemberProfile) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.GuildChangeMemberProfile
}

func (m GuildChangeMemberProfile) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GuildChangeMemberProfile) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
