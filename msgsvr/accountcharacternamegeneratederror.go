package msgsvr

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/d1encoding"
)

type AccountCharacterNameGeneratedError struct {
	Reason int
}

func (m AccountCharacterNameGeneratedError) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.AccountCharacterNameGeneratedError
}

func (m AccountCharacterNameGeneratedError) Serialized() (string, error) {
	return fmt.Sprintf("%d", m.Reason), nil
}

func (m *AccountCharacterNameGeneratedError) Deserialize(extra string) error {
	reason, err := strconv.ParseInt(extra, 10, 32)
	if err != nil {
		return err
	}
	m.Reason = int(reason)

	return nil
}
