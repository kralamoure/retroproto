package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/d1proto"
)

type ExchangeBigStoreType struct {
	ItemType int
}

func (m ExchangeBigStoreType) ProtocolId() d1proto.MsgCliId {
	return d1proto.ExchangeBigStoreType
}

func (m ExchangeBigStoreType) Serialized() (string, error) {
	return fmt.Sprintf("%d", m.ItemType), nil
}

func (m *ExchangeBigStoreType) Deserialize(extra string) error {
	v, err := strconv.ParseInt(extra, 10, 32)
	if err != nil {
		return err
	}
	m.ItemType = int(v)
	return nil
}
