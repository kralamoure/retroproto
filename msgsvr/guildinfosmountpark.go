// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GuildInfosMountPark struct{}

func (m GuildInfosMountPark) MessageId() retroproto.MsgSvrId {
	return retroproto.GuildInfosMountPark
}

func (m GuildInfosMountPark) MessageName() string {
	return "GuildInfosMountPark"
}

func (m GuildInfosMountPark) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildInfosMountPark) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
