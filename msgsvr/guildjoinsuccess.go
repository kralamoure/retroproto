// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GuildJoinSuccess struct{}

func (m GuildJoinSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.GuildJoinSuccess
}

func (m GuildJoinSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildJoinSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
