// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeKillMountInPark struct{}

func (m ExchangeKillMountInPark) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.ExchangeKillMountInPark
}

func (m ExchangeKillMountInPark) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ExchangeKillMountInPark) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
