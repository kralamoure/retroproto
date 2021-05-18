// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GuildLeave struct{}

func (m GuildLeave) ProtocolId() retroproto.MsgSvrId {
	return retroproto.GuildLeave
}

func (m GuildLeave) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildLeave) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
