// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type DocumentsCreateError struct{}

func (m DocumentsCreateError) MessageId() retroproto.MsgSvrId {
	return retroproto.DocumentsCreateError
}

func (m DocumentsCreateError) MessageName() string {
	return "DocumentsCreateError"
}

func (m DocumentsCreateError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *DocumentsCreateError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
