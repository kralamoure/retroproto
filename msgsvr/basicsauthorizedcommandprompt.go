// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type BasicsAuthorizedCommandPrompt struct{}

func (m BasicsAuthorizedCommandPrompt) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.BasicsAuthorizedCommandPrompt
}

func (m BasicsAuthorizedCommandPrompt) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *BasicsAuthorizedCommandPrompt) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
