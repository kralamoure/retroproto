package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type BasicsAuthorizedCommandError struct{}

func (m BasicsAuthorizedCommandError) MessageId() retroproto.MsgSvrId {
	return retroproto.BasicsAuthorizedCommandError
}

func (m BasicsAuthorizedCommandError) Serialized() (string, error) {
	return "", nil
}

func (m *BasicsAuthorizedCommandError) Deserialize(extra string) error {
	return nil
}
