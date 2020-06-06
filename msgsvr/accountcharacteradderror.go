package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type AccountCharacterAddError struct {
	Reason rune
}

func (m AccountCharacterAddError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.AccountCharacterAddError
}

func (m AccountCharacterAddError) Serialized() (string, error) {
	return string(m.Reason), nil
}

func (m *AccountCharacterAddError) Deserialize(extra string) error {
	if len(extra) < 1 {
		return d1proto.ErrInvalidMsg
	}

	m.Reason = rune(extra[0])

	return nil
}
