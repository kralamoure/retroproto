// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GuildCreateSuccess struct{}

func (m GuildCreateSuccess) ProtocolId() retroproto.MsgSvrId {
	return retroproto.GuildCreateSuccess
}

func (m GuildCreateSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildCreateSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
