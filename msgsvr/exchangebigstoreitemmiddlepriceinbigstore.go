package msgsvr

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/retroproto"
)

type ExchangeBigStoreItemMiddlePriceInBigStore struct {
	TemplateId int
	Price      int
}

func (m ExchangeBigStoreItemMiddlePriceInBigStore) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeBigStoreItemMiddlePriceInBigStore
}

func (m ExchangeBigStoreItemMiddlePriceInBigStore) Serialized() (string, error) {
	return fmt.Sprintf("%d|%d", m.TemplateId, m.Price), nil
}

func (m *ExchangeBigStoreItemMiddlePriceInBigStore) Deserialize(extra string) error {
	sli := strings.Split(extra, "|")
	if len(sli) != 2 {
		return retroproto.ErrInvalidMsg
	}

	templateId, err := strconv.ParseInt(sli[0], 10, 32)
	if err != nil {
		return err
	}
	m.TemplateId = int(templateId)

	price, err := strconv.ParseInt(sli[1], 10, 32)
	if err != nil {
		return err
	}
	m.Price = int(price)

	return nil
}
