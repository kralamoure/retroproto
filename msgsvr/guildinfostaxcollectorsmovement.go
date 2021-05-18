// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GuildInfosTaxCollectorsMovement struct{}

func (m GuildInfosTaxCollectorsMovement) ProtocolId() retroproto.MsgSvrId {
	return retroproto.GuildInfosTaxCollectorsMovement
}

func (m GuildInfosTaxCollectorsMovement) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildInfosTaxCollectorsMovement) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
