// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type BasicsWhoIsError struct{}

func (m BasicsWhoIsError) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.BasicsWhoIsError
}

func (m BasicsWhoIsError) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *BasicsWhoIsError) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
