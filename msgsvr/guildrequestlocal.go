// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GuildRequestLocal struct{}

func (m GuildRequestLocal) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GuildRequestLocal
}

func (m GuildRequestLocal) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *GuildRequestLocal) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
