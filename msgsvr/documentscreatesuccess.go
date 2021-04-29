// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type DocumentsCreateSuccess struct{}

func (m DocumentsCreateSuccess) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.DocumentsCreateSuccess
}

func (m DocumentsCreateSuccess) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *DocumentsCreateSuccess) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
