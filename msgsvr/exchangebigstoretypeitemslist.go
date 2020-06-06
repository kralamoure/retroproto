// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeBigStoreTypeItemsList struct{}

func (m ExchangeBigStoreTypeItemsList) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangeBigStoreTypeItemsList
}

func (m ExchangeBigStoreTypeItemsList) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeBigStoreTypeItemsList) Deserialize(extra string) error {
	return nil
}
