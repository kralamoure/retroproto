package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/d1encoding"
)

type ExchangeBigStoreItemList struct {
	ItemTemplateId int
}

func (m ExchangeBigStoreItemList) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.ExchangeBigStoreItemList
}

func (m ExchangeBigStoreItemList) Serialized() (string, error) {
	return fmt.Sprintf("%d", m.ItemTemplateId), nil
}

func (m *ExchangeBigStoreItemList) Deserialize(extra string) error {
	itemTemplateId, err := strconv.ParseInt(extra, 10, 32)
	if err != nil {
		return err
	}
	m.ItemTemplateId = int(itemTemplateId)
	return nil
}
