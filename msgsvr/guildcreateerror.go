// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GuildCreateError struct{}

func (m GuildCreateError) ProtocolId() retroproto.MsgSvrId {
	return retroproto.GuildCreateError
}

func (m GuildCreateError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildCreateError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
