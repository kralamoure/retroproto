// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ItemsFeed struct{}

func NewItemsFeed(extra string) (ItemsFeed, error) {
	var m ItemsFeed

	if err := m.Deserialize(extra); err != nil {
		return ItemsFeed{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ItemsFeed) MessageId() retroproto.MsgCliId {
	return retroproto.ItemsFeed
}

func (m ItemsFeed) MessageName() string {
	return "ItemsFeed"
}

func (m ItemsFeed) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ItemsFeed) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
