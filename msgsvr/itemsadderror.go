package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ItemsAddError struct {
	Reason rune
}

func NewItemsAddError(extra string) (ItemsAddError, error) {
	var m ItemsAddError

	if err := m.Deserialize(extra); err != nil {
		return ItemsAddError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ItemsAddError) MessageId() retroproto.MsgSvrId {
	return retroproto.ItemsAddError
}

func (m ItemsAddError) MessageName() string {
	return "ItemsAddError"
}

func (m ItemsAddError) Serialized() (string, error) {
	return string(m.Reason), nil
}

func (m *ItemsAddError) Deserialize(extra string) error {
	if len(extra) == 0 {
		return retroproto.ErrInvalidMsg
	}

	for _, v := range extra {
		m.Reason = v
		break
	}

	return nil
}
