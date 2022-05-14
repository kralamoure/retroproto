package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retroproto"
)

type ExchangeGetItemMiddlePriceInBigStore struct {
	TemplateId int
}

func NewExchangeGetItemMiddlePriceInBigStore(extra string) (ExchangeGetItemMiddlePriceInBigStore, error) {
	var m ExchangeGetItemMiddlePriceInBigStore

	if err := m.Deserialize(extra); err != nil {
		return ExchangeGetItemMiddlePriceInBigStore{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeGetItemMiddlePriceInBigStore) MessageId() retroproto.MsgCliId {
	return retroproto.ExchangeGetItemMiddlePriceInBigStore
}

func (m ExchangeGetItemMiddlePriceInBigStore) MessageName() string {
	return "ExchangeGetItemMiddlePriceInBigStore"
}

func (m ExchangeGetItemMiddlePriceInBigStore) Serialized() (string, error) {
	return fmt.Sprintf("%d", m.TemplateId), nil
}

func (m *ExchangeGetItemMiddlePriceInBigStore) Deserialize(extra string) error {
	templateId, err := strconv.ParseInt(extra, 10, 32)
	if err != nil {
		return err
	}
	m.TemplateId = int(templateId)

	return nil
}
