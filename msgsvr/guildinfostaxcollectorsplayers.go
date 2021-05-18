// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GuildInfosTaxCollectorsPlayers struct{}

func (m GuildInfosTaxCollectorsPlayers) ProtocolId() retroproto.MsgSvrId {
	return retroproto.GuildInfosTaxCollectorsPlayers
}

func (m GuildInfosTaxCollectorsPlayers) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildInfosTaxCollectorsPlayers) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
