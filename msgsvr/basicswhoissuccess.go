// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type BasicsWhoIsSuccess struct{}

func (m BasicsWhoIsSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.BasicsWhoIsSuccess
}

func (m BasicsWhoIsSuccess) MessageName() string {
	return "BasicsWhoIsSuccess"
}

func (m BasicsWhoIsSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *BasicsWhoIsSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
