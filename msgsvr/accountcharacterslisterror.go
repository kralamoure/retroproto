// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type AccountCharactersListError struct{}

func (m AccountCharactersListError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.AccountCharactersListError
}

func (m AccountCharactersListError) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *AccountCharactersListError) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
