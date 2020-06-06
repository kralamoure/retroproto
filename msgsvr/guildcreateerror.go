// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GuildCreateError struct{}

func (m GuildCreateError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GuildCreateError
}

func (m GuildCreateError) Serialized() (string, error) {
	return "", nil
}

func (m *GuildCreateError) Deserialize(extra string) error {
	return nil
}
