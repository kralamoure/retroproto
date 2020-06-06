// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type GuildRequestLeave struct{}

func (m GuildRequestLeave) ProtocolId() d1proto.MsgCliId {
	return d1proto.GuildRequestLeave
}

func (m GuildRequestLeave) Serialized() (string, error) {
	return "", nil
}

func (m *GuildRequestLeave) Deserialize(extra string) error {
	return nil
}
