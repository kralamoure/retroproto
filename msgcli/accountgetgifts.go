package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type AccountGetGifts struct {
	Lang string
}

func (m AccountGetGifts) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.AccountGetGifts
}

func (m AccountGetGifts) Serialized() (string, error) {
	return m.Lang, nil
}

func (m *AccountGetGifts) Deserialize(extra string) error {
	m.Lang = extra

	return nil
}
