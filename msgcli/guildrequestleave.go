// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type GuildRequestLeave struct{}

func (m GuildRequestLeave) MessageId() retroproto.MsgCliId {
	return retroproto.GuildRequestLeave
}

func (m GuildRequestLeave) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildRequestLeave) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
