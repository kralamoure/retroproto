package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retroproto"
)

type ExchangeBigStoreItemList struct {
	ItemTemplateId int
}

func NewExchangeBigStoreItemList(extra string) (ExchangeBigStoreItemList, error) {
	var m ExchangeBigStoreItemList

	if err := m.Deserialize(extra); err != nil {
		return ExchangeBigStoreItemList{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeBigStoreItemList) MessageId() retroproto.MsgCliId {
	return retroproto.ExchangeBigStoreItemList
}

func (m ExchangeBigStoreItemList) MessageName() string {
	return "ExchangeBigStoreItemList"
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
