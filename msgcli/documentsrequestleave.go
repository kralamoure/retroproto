// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type DocumentsRequestLeave struct{}

func (m DocumentsRequestLeave) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.DocumentsRequestLeave
}

func (m DocumentsRequestLeave) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *DocumentsRequestLeave) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
