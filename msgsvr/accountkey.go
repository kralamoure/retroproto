// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AccountKey struct{}

func NewAccountKey(extra string) (AccountKey, error) {
	var m AccountKey

	if err := m.Deserialize(extra); err != nil {
		return AccountKey{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountKey) MessageId() retroproto.MsgSvrId {
	return retroproto.AccountKey
}

func (m AccountKey) MessageName() string {
	return "AccountKey"
}

func (m AccountKey) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *AccountKey) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
