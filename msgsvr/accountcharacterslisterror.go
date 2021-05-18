// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type AccountCharactersListError struct{}

func (m AccountCharactersListError) ProtocolId() retroproto.MsgSvrId {
	return retroproto.AccountCharactersListError
}

func (m AccountCharactersListError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *AccountCharactersListError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
