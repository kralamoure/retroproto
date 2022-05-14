package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type AccountGetGifts struct {
	Lang string
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
