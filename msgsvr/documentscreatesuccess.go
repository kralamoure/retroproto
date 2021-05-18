// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type DocumentsCreateSuccess struct{}

func (m DocumentsCreateSuccess) ProtocolId() retroproto.MsgSvrId {
	return retroproto.DocumentsCreateSuccess
}

func (m DocumentsCreateSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *DocumentsCreateSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
