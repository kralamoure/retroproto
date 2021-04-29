// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GuildJoinError struct{}

func (m GuildJoinError) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GuildJoinError
}

func (m GuildJoinError) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GuildJoinError) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
