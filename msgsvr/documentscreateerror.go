// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type DocumentsCreateError struct{}

func (m DocumentsCreateError) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.DocumentsCreateError
}

func (m DocumentsCreateError) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *DocumentsCreateError) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
