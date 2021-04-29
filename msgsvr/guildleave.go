// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GuildLeave struct{}

func (m GuildLeave) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GuildLeave
}

func (m GuildLeave) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GuildLeave) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
