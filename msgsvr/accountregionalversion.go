package msgsvr

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/d1proto"
)

type AccountRegionalVersion struct {
	Value int
}

func (m AccountRegionalVersion) ProtocolId() d1proto.MsgSvrId {
	return d1proto.AccountRegionalVersion
}

func (m AccountRegionalVersion) Serialized() (string, error) {
	return fmt.Sprintf("%d", m.Value), nil
}

func (m *AccountRegionalVersion) Deserialize(extra string) error {
	if len(extra) < 1 {
		return d1proto.ErrInvalidMsg
	}

	value, err := strconv.ParseInt(extra, 10, 32)
	if err != nil {
		return err
	}
	m.Value = int(value)

	return nil
}
