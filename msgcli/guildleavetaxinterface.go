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
	return "", d1proto.ErrNotImplemented
}

func (m *GuildLeaveTaxInterface) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
