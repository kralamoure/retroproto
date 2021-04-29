// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GuildInfosTaxCollectorsMovement struct{}

func (m GuildInfosTaxCollectorsMovement) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GuildInfosTaxCollectorsMovement
}

func (m GuildInfosTaxCollectorsMovement) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GuildInfosTaxCollectorsMovement) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
