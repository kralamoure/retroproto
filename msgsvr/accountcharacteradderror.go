package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type AccountCharacterAddError struct {
	Reason rune
}

func (m AccountCharacterAddError) ProtocolId() retroproto.MsgSvrId {
	return retroproto.AccountCharacterAddError
}

func (m AccountCharacterAddError) Serialized() (string, error) {
	return string(m.Reason), nil
}

func (m *AccountCharacterAddError) Deserialize(extra string) error {
	if len(extra) < 1 {
		return retroproto.ErrInvalidMsg
	}

	m.Reason = rune(extra[0])

	return nil
}
