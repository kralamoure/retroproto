// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GuildRequestDistant struct{}

func (m GuildRequestDistant) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GuildRequestDistant
}

func (m GuildRequestDistant) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GuildRequestDistant) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
