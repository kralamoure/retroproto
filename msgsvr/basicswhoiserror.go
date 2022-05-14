// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type BasicsWhoIsError struct{}

func (m BasicsWhoIsError) MessageId() retroproto.MsgSvrId {
	return retroproto.BasicsWhoIsError
}

func (m BasicsWhoIsError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *BasicsWhoIsError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
