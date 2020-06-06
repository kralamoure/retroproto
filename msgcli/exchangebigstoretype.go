// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeBigStoreType struct{}

func (m ExchangeBigStoreType) ProtocolId() d1proto.MsgCliId {
	return d1proto.ExchangeBigStoreType
}

func (m ExchangeBigStoreType) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeBigStoreType) Deserialize(extra string) error {
	return nil
}
