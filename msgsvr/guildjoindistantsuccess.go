// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GuildJoinDistantSuccess struct{}

func (m GuildJoinDistantSuccess) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GuildJoinDistantSuccess
}

func (m GuildJoinDistantSuccess) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GuildJoinDistantSuccess) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
