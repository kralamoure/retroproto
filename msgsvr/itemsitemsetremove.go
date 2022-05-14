package msgsvr

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/retroproto"
)

type ItemsItemSetRemove struct {
	Id int
}

func NewItemsItemSetRemove(extra string) (ItemsItemSetRemove, error) {
	var m ItemsItemSetRemove

	if err := m.Deserialize(extra); err != nil {
		return ItemsItemSetRemove{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ItemsItemSetRemove) MessageId() retroproto.MsgSvrId {
	return retroproto.ItemsItemSetRemove
}

func (m ItemsItemSetRemove) MessageName() string {
	return "ItemsItemSetRemove"
}

func (m ItemsItemSetRemove) Serialized() (string, error) {
	return fmt.Sprintf("%d", m.Id), nil
}

func (m *ItemsItemSetRemove) Deserialize(extra string) error {
	sli := strings.Split(extra, "|")

	if sli[0] != "" {
		id, err := strconv.ParseInt(sli[0], 10, 32)
		if err != nil {
			return err
		}
		m.Id = int(id)
	}

	return nil
}
