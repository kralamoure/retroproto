package typ

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/d1proto"
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
	return fmt.Sprintf("%d;%s;%d;%d;%d", m.Id, effects, m.PriceSet1, m.PriceSet2, m.PriceSet3), nil
}

func (m *ExchangeBigStoreItemsListItem) Deserialize(extra string) error {
	sli := strings.Split(extra, ";")
	if len(sli) != 5 {
		return d1proto.ErrInvalidMsg
	}

	id, err := strconv.ParseInt(sli[0], 10, 32)
	if err != nil {
		return err
	}
	m.Id = int(id)

	if sli[1] != "" {
		m.Effects = strings.Split(sli[1], ",")
	}

	priceSet1, err := strconv.ParseInt(sli[2], 10, 32)
	if err != nil {
		return err
	}
	m.PriceSet1 = int(priceSet1)

	priceSet2, err := strconv.ParseInt(sli[3], 10, 32)
	if err != nil {
		return err
	}
	m.PriceSet2 = int(priceSet2)

	priceSet3, err := strconv.ParseInt(sli[4], 10, 32)
	if err != nil {
		return err
	}
	m.PriceSet3 = int(priceSet3)

	return nil
}
