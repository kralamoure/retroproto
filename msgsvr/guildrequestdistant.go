// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GuildRequestDistant struct{}

func (m GuildRequestDistant) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GuildRequestDistant
}

func (m GuildRequestDistant) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *GuildRequestDistant) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
