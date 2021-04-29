// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type AccountCharacterDeleteSuccess struct{}

func (m AccountCharacterDeleteSuccess) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.AccountCharacterDeleteSuccess
}

func (m AccountCharacterDeleteSuccess) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *AccountCharacterDeleteSuccess) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
