// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type GuildLeaveTaxInterface struct{}

func (m GuildLeaveTaxInterface) ProtocolId() d1proto.MsgCliId {
	return d1proto.GuildLeaveTaxInterface
}

func (m GuildLeaveTaxInterface) Serialized() (string, error) {
	return "", nil
}

func (m *GuildLeaveTaxInterface) Deserialize(extra string) error {
	return nil
}
