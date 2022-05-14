// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type DocumentsLeave struct{}

func (m DocumentsLeave) MessageId() retroproto.MsgSvrId {
	return retroproto.DocumentsLeave
}

func (m DocumentsLeave) MessageName() string {
	return "DocumentsLeave"
}

func (m DocumentsLeave) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *DocumentsLeave) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
