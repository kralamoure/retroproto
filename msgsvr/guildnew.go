// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GuildNew struct{}

func (m GuildNew) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GuildNew
}

func (m GuildNew) Serialized() (string, error) {
	return "", nil
}

func (m *GuildNew) Deserialize(extra string) error {
	return nil
}
