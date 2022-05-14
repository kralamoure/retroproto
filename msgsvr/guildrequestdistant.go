// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GuildRequestDistant struct{}

func (m GuildRequestDistant) MessageId() retroproto.MsgSvrId {
	return retroproto.GuildRequestDistant
}

func (m GuildRequestDistant) MessageName() string {
	return "GuildRequestDistant"
}

func (m GuildRequestDistant) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildRequestDistant) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
