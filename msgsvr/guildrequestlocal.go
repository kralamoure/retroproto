// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GuildRequestLocal struct{}

func (m GuildRequestLocal) MessageId() retroproto.MsgSvrId {
	return retroproto.GuildRequestLocal
}

func (m GuildRequestLocal) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildRequestLocal) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
