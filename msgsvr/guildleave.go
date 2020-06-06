// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GuildLeave struct{}

func (m GuildLeave) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GuildLeave
}

func (m GuildLeave) Serialized() (string, error) {
	return "", nil
}

func (m *GuildLeave) Deserialize(extra string) error {
	return nil
}
