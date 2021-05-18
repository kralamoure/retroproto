// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type BasicsWhoIsSuccess struct{}

func (m BasicsWhoIsSuccess) ProtocolId() retroproto.MsgSvrId {
	return retroproto.BasicsWhoIsSuccess
}

func (m BasicsWhoIsSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *BasicsWhoIsSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
