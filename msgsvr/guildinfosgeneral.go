// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GuildInfosGeneral struct{}

func (m GuildInfosGeneral) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GuildInfosGeneral
}

func (m GuildInfosGeneral) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GuildInfosGeneral) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
