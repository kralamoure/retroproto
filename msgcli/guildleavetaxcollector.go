// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type GuildLeaveTaxCollector struct{}

func (m GuildLeaveTaxCollector) ProtocolId() d1proto.MsgCliId {
	return d1proto.GuildLeaveTaxCollector
}

func (m GuildLeaveTaxCollector) Serialized() (string, error) {
	return "", nil
}

func (m *GuildLeaveTaxCollector) Deserialize(extra string) error {
	return nil
}
