package msgsvr

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retroproto"
)

type AccountRegionalVersion struct {
	Value int
}

func NewAccountRegionalVersion(extra string) (AccountRegionalVersion, error) {
	var m AccountRegionalVersion

	if err := m.Deserialize(extra); err != nil {
		return AccountRegionalVersion{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountRegionalVersion) MessageId() retroproto.MsgSvrId {
	return retroproto.AccountRegionalVersion
}

func (m AccountRegionalVersion) MessageName() string {
	return "AccountRegionalVersion"
}

func (m AccountRegionalVersion) Serialized() (string, error) {
	return fmt.Sprintf("%d", m.Value), nil
}

func (m *AccountRegionalVersion) Deserialize(extra string) error {
	if len(extra) < 1 {
		return retroproto.ErrInvalidMsg
	}

	value, err := strconv.ParseInt(extra, 10, 32)
	if err != nil {
		return err
	}
	m.Value = int(value)

	return nil
}
