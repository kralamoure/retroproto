// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type AccountCharacterDeleteSuccess struct{}

func (m AccountCharacterDeleteSuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.AccountCharacterDeleteSuccess
}

func (m AccountCharacterDeleteSuccess) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *AccountCharacterDeleteSuccess) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
