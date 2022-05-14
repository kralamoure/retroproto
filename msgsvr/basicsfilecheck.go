// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type BasicsFileCheck struct{}

func (m BasicsFileCheck) MessageId() retroproto.MsgSvrId {
	return retroproto.BasicsFileCheck
}

func (m BasicsFileCheck) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *BasicsFileCheck) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
