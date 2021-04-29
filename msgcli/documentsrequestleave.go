// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type DocumentsRequestLeave struct{}

func (m DocumentsRequestLeave) ProtocolId() d1proto.MsgCliId {
	return d1proto.DocumentsRequestLeave
}

func (m DocumentsRequestLeave) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *DocumentsRequestLeave) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
