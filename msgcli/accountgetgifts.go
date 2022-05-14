package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AccountGetGifts struct {
	Lang string
}

func NewAccountGetGifts(extra string) (AccountGetGifts, error) {
	var m AccountGetGifts

	if err := m.Deserialize(extra); err != nil {
		return AccountGetGifts{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountGetGifts) MessageId() retroproto.MsgCliId {
	return retroproto.AccountGetGifts
}

func (m AccountGetGifts) MessageName() string {
	return "AccountGetGifts"
}

func (m AccountGetGifts) Serialized() (string, error) {
	return m.Lang, nil
}

func (m *AccountGetGifts) Deserialize(extra string) error {
	m.Lang = extra

	return nil
}
