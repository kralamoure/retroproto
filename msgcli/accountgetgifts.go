package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type AccountGetGifts struct {
	Lang string
}

func (m AccountGetGifts) ProtocolId() d1proto.MsgCliId {
	return d1proto.AccountGetGifts
}

func (m AccountGetGifts) Serialized() (string, error) {
	return m.Lang, nil
}

func (m *AccountGetGifts) Deserialize(extra string) error {
	m.Lang = extra

	return nil
}
