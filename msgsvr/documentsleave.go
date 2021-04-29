// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type DocumentsLeave struct{}

func (m DocumentsLeave) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.DocumentsLeave
}

func (m DocumentsLeave) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *DocumentsLeave) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
