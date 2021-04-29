// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type DocumentsLeave struct{}

func (m DocumentsLeave) ProtocolId() d1proto.MsgSvrId {
	return d1proto.DocumentsLeave
}

func (m DocumentsLeave) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *DocumentsLeave) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
