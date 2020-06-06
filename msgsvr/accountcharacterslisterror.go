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
	return "", nil
}

func (m *AccountCharactersListError) Deserialize(extra string) error {
	return nil
}
