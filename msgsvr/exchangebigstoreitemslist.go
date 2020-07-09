// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeBigStoreItemsList struct{}

func (m ExchangeBigStoreItemsList) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangeBigStoreItemsList
}

func (m ExchangeBigStoreItemsList) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeBigStoreItemsList) Deserialize(extra string) error {
	return nil
}
