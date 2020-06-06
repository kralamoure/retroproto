// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeBigStoreItemMiddlePriceInBigStore struct{}

func (m ExchangeBigStoreItemMiddlePriceInBigStore) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangeBigStoreItemMiddlePriceInBigStore
}

func (m ExchangeBigStoreItemMiddlePriceInBigStore) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeBigStoreItemMiddlePriceInBigStore) Deserialize(extra string) error {
	return nil
}
