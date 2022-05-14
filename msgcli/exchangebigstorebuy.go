package msgcli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/retroproto"
)

type ExchangeBigStoreBuy struct {
	ItemId        int
	QuantityIndex int
	Price         int
}

func (m ExchangeBigStoreBuy) MessageId() retroproto.MsgCliId {
	return retroproto.ExchangeBigStoreBuy
}

func (m ExchangeBigStoreBuy) MessageName() string {
	return "ExchangeBigStoreBuy"
}

func (m ExchangeBigStoreBuy) Serialized() (string, error) {
	return fmt.Sprintf("%d|%d|%d", m.ItemId, m.QuantityIndex, m.Price), nil
}

func (m *ExchangeBigStoreBuy) Deserialize(extra string) error {
	sli := strings.Split(extra, "|")
	if len(sli) != 3 {
		return retroproto.ErrInvalidMsg
	}

	itemId, err := strconv.ParseInt(sli[0], 10, 32)
	if err != nil {
		return err
	}
	m.ItemId = int(itemId)

	quantityIndex, err := strconv.ParseInt(sli[1], 10, 32)
	if err != nil {
		return err
	}
	m.QuantityIndex = int(quantityIndex)

	price, err := strconv.ParseInt(sli[2], 10, 32)
	if err != nil {
		return err
	}
	m.Price = int(price)

	return nil
}
