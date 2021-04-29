package typ

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/d1encoding"
)

type ExchangeBigStoreItemsListItem struct {
	Id        int
	Effects   []string
	PriceSet1 int
	PriceSet2 int
	PriceSet3 int
}

func (m ExchangeBigStoreItemsListItem) Serialized() (string, error) {
	effects := strings.Join(m.Effects, ",")

	var priceSet1 string
	if m.PriceSet1 > 0 {
		priceSet1 = fmt.Sprintf("%d", m.PriceSet1)
	}

	var priceSet2 string
	if m.PriceSet2 > 0 {
		priceSet2 = fmt.Sprintf("%d", m.PriceSet2)
	}

	var priceSet3 string
	if m.PriceSet3 > 0 {
		priceSet3 = fmt.Sprintf("%d", m.PriceSet3)
	}

	return fmt.Sprintf("%d;%s;%s;%s;%s", m.Id, effects, priceSet1, priceSet2, priceSet3), nil
}

func (m *ExchangeBigStoreItemsListItem) Deserialize(extra string) error {
	sli := strings.Split(extra, ";")
	if len(sli) != 5 {
		return d1encoding.ErrInvalidMsg
	}

	id, err := strconv.ParseInt(sli[0], 10, 32)
	if err != nil {
		return err
	}
	m.Id = int(id)

	if sli[1] != "" {
		m.Effects = strings.Split(sli[1], ",")
	}

	if sli[2] != "" {
		priceSet1, err := strconv.ParseInt(sli[2], 10, 32)
		if err != nil {
			return err
		}
		m.PriceSet1 = int(priceSet1)
	}

	if sli[3] != "" {
		priceSet2, err := strconv.ParseInt(sli[3], 10, 32)
		if err != nil {
			return err
		}
		m.PriceSet2 = int(priceSet2)
	}

	if sli[4] != "" {
		priceSet3, err := strconv.ParseInt(sli[4], 10, 32)
		if err != nil {
			return err
		}
		m.PriceSet3 = int(priceSet3)
	}

	return nil
}
