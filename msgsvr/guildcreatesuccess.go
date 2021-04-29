// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GuildCreateSuccess struct{}

func (m GuildCreateSuccess) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GuildCreateSuccess
}

func (m GuildCreateSuccess) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GuildCreateSuccess) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
