// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeKillMountInPark struct{}

func (m ExchangeKillMountInPark) ProtocolId() d1proto.MsgCliId {
	return d1proto.ExchangeKillMountInPark
}

func (m ExchangeKillMountInPark) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ExchangeKillMountInPark) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
