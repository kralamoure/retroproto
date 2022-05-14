// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type BasicsAuthorizedCommandPrompt struct{}

func (m BasicsAuthorizedCommandPrompt) MessageId() retroproto.MsgSvrId {
	return retroproto.BasicsAuthorizedCommandPrompt
}

func (m BasicsAuthorizedCommandPrompt) MessageName() string {
	return "BasicsAuthorizedCommandPrompt"
}

func (m BasicsAuthorizedCommandPrompt) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *BasicsAuthorizedCommandPrompt) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
