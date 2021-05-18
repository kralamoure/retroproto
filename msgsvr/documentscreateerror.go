// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type DocumentsCreateError struct{}

func (m DocumentsCreateError) ProtocolId() retroproto.MsgSvrId {
	return retroproto.DocumentsCreateError
}

func (m DocumentsCreateError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *DocumentsCreateError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
