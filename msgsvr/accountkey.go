// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type AccountKey struct{}

func (m AccountKey) ProtocolId() d1proto.MsgSvrId {
	return d1proto.AccountKey
}

func (m AccountKey) Serialized() (string, error) {
	return "", nil
}

func (m *AccountKey) Deserialize(extra string) error {
	return nil
}
