// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type DocumentsCreateSuccess struct{}

func (m DocumentsCreateSuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.DocumentsCreateSuccess
}

func (m DocumentsCreateSuccess) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *DocumentsCreateSuccess) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
