package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type AccountCharacterAddError struct {
	Reason rune
}

func (m AccountCharacterAddError) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.AccountCharacterAddError
}

func (m AccountCharacterAddError) Serialized() (string, error) {
	return string(m.Reason), nil
}

func (m *AccountCharacterAddError) Deserialize(extra string) error {
	if len(extra) < 1 {
		return d1encoding.ErrInvalidMsg
	}

	m.Reason = rune(extra[0])

	return nil
}
