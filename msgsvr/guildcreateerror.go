// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GuildCreateError struct{}

func (m GuildCreateError) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GuildCreateError
}

func (m GuildCreateError) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GuildCreateError) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
