// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type AccountCharactersListError struct{}

func (m AccountCharactersListError) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.AccountCharactersListError
}

func (m AccountCharactersListError) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *AccountCharactersListError) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
