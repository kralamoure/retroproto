package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/d1proto"
)

type ExchangeGetItemMiddlePriceInBigStore struct {
	TemplateId int
}

func (m ExchangeGetItemMiddlePriceInBigStore) ProtocolId() d1proto.MsgCliId {
	return d1proto.ExchangeGetItemMiddlePriceInBigStore
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
