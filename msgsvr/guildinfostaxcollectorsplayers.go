// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GuildInfosTaxCollectorsPlayers struct{}

func (m GuildInfosTaxCollectorsPlayers) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GuildInfosTaxCollectorsPlayers
}

func (m GuildInfosTaxCollectorsPlayers) Serialized() (string, error) {
	return "", nil
}

func (m *GuildInfosTaxCollectorsPlayers) Deserialize(extra string) error {
	return nil
}
