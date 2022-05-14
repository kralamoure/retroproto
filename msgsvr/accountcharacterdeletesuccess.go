// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type AccountCharacterDeleteSuccess struct{}

func (m AccountCharacterDeleteSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.AccountCharacterDeleteSuccess
}

func (m AccountCharacterDeleteSuccess) MessageName() string {
	return "AccountCharacterDeleteSuccess"
}

func (m AccountCharacterDeleteSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *AccountCharacterDeleteSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
