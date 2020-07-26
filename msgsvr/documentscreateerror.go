// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type DocumentsCreateError struct{}

func (m DocumentsCreateError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.DocumentsCreateError
}

func (m DocumentsCreateError) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *DocumentsCreateError) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
