// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GuildInfosMountPark struct{}

func (m GuildInfosMountPark) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GuildInfosMountPark
}

func (m GuildInfosMountPark) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GuildInfosMountPark) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
