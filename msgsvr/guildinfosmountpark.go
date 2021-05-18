// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GuildInfosMountPark struct{}

func (m GuildInfosMountPark) ProtocolId() retroproto.MsgSvrId {
	return retroproto.GuildInfosMountPark
}

func (m GuildInfosMountPark) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildInfosMountPark) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
