// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type BasicsWhoIsSuccess struct{}

func (m BasicsWhoIsSuccess) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.BasicsWhoIsSuccess
}

func (m BasicsWhoIsSuccess) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *BasicsWhoIsSuccess) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
