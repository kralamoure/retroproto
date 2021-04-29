// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GuildNew struct{}

func (m GuildNew) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GuildNew
}

func (m GuildNew) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GuildNew) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
