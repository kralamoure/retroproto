// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type GuildRequestLeave struct{}

func (m GuildRequestLeave) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.GuildRequestLeave
}

func (m GuildRequestLeave) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GuildRequestLeave) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
