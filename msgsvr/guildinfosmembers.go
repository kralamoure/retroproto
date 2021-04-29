// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GuildInfosMembers struct{}

func (m GuildInfosMembers) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GuildInfosMembers
}

func (m GuildInfosMembers) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GuildInfosMembers) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
