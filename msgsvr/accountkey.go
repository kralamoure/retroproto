// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type AccountKey struct{}

func (m AccountKey) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.AccountKey
}

func (m AccountKey) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *AccountKey) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
