// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GuildRequestLocal struct{}

func (m GuildRequestLocal) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GuildRequestLocal
}

func (m GuildRequestLocal) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GuildRequestLocal) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
