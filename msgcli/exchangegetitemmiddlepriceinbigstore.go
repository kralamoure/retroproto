// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeGetItemMiddlePriceInBigStore struct{}

func (m ExchangeGetItemMiddlePriceInBigStore) ProtocolId() d1proto.MsgCliId {
	return d1proto.ExchangeGetItemMiddlePriceInBigStore
}

func (m ExchangeGetItemMiddlePriceInBigStore) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeGetItemMiddlePriceInBigStore) Deserialize(extra string) error {
	return nil
}
