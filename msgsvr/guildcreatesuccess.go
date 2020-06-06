// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GuildCreateSuccess struct{}

func (m GuildCreateSuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GuildCreateSuccess
}

func (m GuildCreateSuccess) Serialized() (string, error) {
	return "", nil
}

func (m *GuildCreateSuccess) Deserialize(extra string) error {
	return nil
}
