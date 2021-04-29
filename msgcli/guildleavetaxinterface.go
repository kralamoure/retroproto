// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type GuildLeaveTaxInterface struct{}

func (m GuildLeaveTaxInterface) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.GuildLeaveTaxInterface
}

func (m GuildLeaveTaxInterface) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GuildLeaveTaxInterface) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
