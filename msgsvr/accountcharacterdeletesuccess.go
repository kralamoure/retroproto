// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type AccountCharacterDeleteSuccess struct{}

func (m AccountCharacterDeleteSuccess) ProtocolId() retroproto.MsgSvrId {
	return retroproto.AccountCharacterDeleteSuccess
}

func (m AccountCharacterDeleteSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *AccountCharacterDeleteSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
