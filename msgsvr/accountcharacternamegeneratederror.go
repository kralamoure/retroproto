package msgsvr

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/d1proto"
)

type AccountCharacterNameGeneratedError struct {
	Reason int
}

func (m AccountCharacterNameGeneratedError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.AccountCharacterNameGeneratedError
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
